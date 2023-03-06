// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compile

import (
	"bytes"
	"context"
	"fmt"
	"hash/crc32"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/cnservice/cnclient"
	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/common/morpc"
	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/fileservice"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/matrixorigin/matrixone/pkg/pb/pipeline"
	"github.com/matrixorigin/matrixone/pkg/pb/plan"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/txn/client"
	"github.com/matrixorigin/matrixone/pkg/vm/engine"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const (
	maxMessageSizeToMoRpc = 256 * mpool.KB
)

// cnInformation records service information to help handle messages.
type cnInformation struct {
	storeEngine engine.Engine
	fileService fileservice.FileService
}

// processHelper records source process information to help
// rebuild the process at the remote node.
type processHelper struct {
	id               string
	lim              process.Limitation
	unixTime         int64
	txnOperator      client.TxnOperator
	txnClient        client.TxnClient
	sessionInfo      process.SessionInfo
	analysisNodeList []int32
}

// messageSenderOnClient is a structure
// for sending message and receiving results on cn-client.
type messageSenderOnClient struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	streamSender morpc.Stream
	receiveCh    chan morpc.Message
}

func newMessageSenderOnClient(
	ctx context.Context, toAddr string) (messageSenderOnClient, error) {
	var sender = messageSenderOnClient{}

	streamSender, err := cnclient.GetStreamSender(toAddr)
	if err != nil {
		return sender, nil
	}

	sender.streamSender = streamSender
	if _, ok := ctx.Deadline(); !ok {
		sender.ctx, sender.ctxCancel = context.WithTimeout(ctx, time.Second*10000)
	} else {
		sender.ctx = ctx
	}
	return sender, nil
}

// XXX we can set a scope as argument directly next day.
func (sender *messageSenderOnClient) send(
	scopeData, procData []byte, messageType uint64) error {
	sdLen := len(scopeData)
	fmt.Printf("[mymylog] sLen = %d\n", sdLen)
	if sdLen <= maxMessageSizeToMoRpc {
		message := cnclient.AcquireMessage()
		message.SetID(sender.streamSender.ID())
		message.SetMessageType(messageType)
		message.SetData(scopeData)
		message.SetProcData(procData)
		message.SetSequence(0)
		message.SetSid(pipeline.Last)
		return sender.streamSender.Send(sender.ctx, message)
	}

	fmt.Printf("[pipeline msg] seperate pipeine then send\n")
	start := 0
	cnt := uint64(0)
	uu := uuid.New()
	for start < sdLen {
		end := start + maxMessageSizeToMoRpc
		isLast := end > sdLen

		message := cnclient.AcquireMessage()
		message.SetID(sender.streamSender.ID())
		message.SetMessageType(messageType)
		message.SetSequence(cnt)
		message.Uuid = uu[:]
		if isLast {
			message.SetData(scopeData[start:sdLen])
			message.SetProcData(procData)
			message.SetSid(pipeline.Last)
		} else {
			message.SetData(scopeData[start:end])
			message.SetProcData(nil)
			message.SetSid(pipeline.WaitingNext)
		}

		if err := sender.streamSender.Send(sender.ctx, message); err != nil {
			return err
		}
		cnt++
		start = end
	}
	fmt.Printf("[pipeline msg] uuid %s seperate pipeine send done. cnt = %d\n", uu, cnt)
	return nil
}

func (sender *messageSenderOnClient) receiveMessage() (morpc.Message, error) {
	var err error
	if sender.receiveCh == nil {
		sender.receiveCh, err = sender.streamSender.Receive()
		if err != nil {
			return nil, err
		}
	}

	select {
	case <-sender.ctx.Done():
		return nil, moerr.NewRPCTimeout(sender.ctx)
	case val, ok := <-sender.receiveCh:
		if !ok || val == nil {
			// ch close
			return nil, moerr.NewStreamClosed(sender.ctx)
		}
		return val, nil
	}
}

func (sender *messageSenderOnClient) close() {
	if sender.ctxCancel != nil {
		sender.ctxCancel()
	}
	// XXX not a good way to deal it if close failed.
	_ = sender.streamSender.Close()
}

// messageReceiverOnServer is a structure
// for processing received message and writing results back at cn-server.
type messageReceiverOnServer struct {
	ctx         context.Context
	messageId   uint64
	messageTyp  uint64
	messageUuid uuid.UUID

	cnInformation cnInformation
	// information to build a process.
	procBuildHelper processHelper

	clientSession   morpc.ClientSession
	messageAcquirer func() morpc.Message
	maxMessageSize  int
	scopeData       []byte

	// XXX what's that. So confused.
	sequence uint64

	// result.
	finalAnalysisInfo []*process.AnalyzeInfo
}

func newMessageReceiverOnServer(
	ctx context.Context,
	m *pipeline.Message,
	cs morpc.ClientSession,
	messageAcquirer func() morpc.Message,
	storeEngine engine.Engine,
	fileService fileservice.FileService,
	txnClient client.TxnClient) messageReceiverOnServer {

	receiver := messageReceiverOnServer{
		ctx:             ctx,
		messageId:       m.GetId(),
		messageTyp:      m.GetCmd(),
		clientSession:   cs,
		messageAcquirer: messageAcquirer,
		maxMessageSize:  maxMessageSizeToMoRpc,
	}

	switch m.GetCmd() {
	case pipeline.PrepareDoneNotifyMessage:
		opUuid, err := uuid.FromBytes(m.GetUuid())
		if err != nil {
			logutil.Errorf("decode uuid from pipeline.Message failed, bytes are %v", m.GetUuid())
			panic("cn receive a message with wrong uuid bytes")
		}
		receiver.messageUuid = opUuid

	case pipeline.PipelineMessage:
		var err error
		receiver.cnInformation = cnInformation{
			storeEngine: storeEngine,
			fileService: fileService,
		}
		receiver.procBuildHelper, err = generateProcessHelper(m.GetProcInfoData(), txnClient)
		if err != nil {
			logutil.Errorf("decode process info from pipeline.Message failed, bytes are %v", m.GetProcInfoData())
			panic("cn receive a message with wrong process bytes")
		}

		completeData, err := getCompleteScopeDate(ctx, int(m.GetSequence()), m, cs)
		if err != nil {
			logutil.Errorf("failed to get complete pipeline data. err = %s", err)
			panic("failed to get complete pipeline data")
		}
		receiver.scopeData = completeData
		// TODO: remove it
		opUuid, _ := uuid.FromBytes(m.GetUuid())
		receiver.messageUuid = opUuid
	default:
		logutil.Errorf("unknown cmd %d for pipeline.Message", m.GetCmd())
		panic("unknown message type")
	}

	return receiver
}

func getCompleteScopeDate(ctx context.Context, msgCount int, m *pipeline.Message, cs morpc.ClientSession) ([]byte, error) {
	if msgCount == 0 { // status == Last && sequence == 0 means it is a complete data
		return m.Data, nil
	}

	msgVec := make([]*pipeline.Message, int(msgCount))
	cache, err := cs.CreateCache(ctx, 0)
	if err != nil {
		return nil, err
	}

	cnt := 0
	for cnt < msgCount {
		msg, ok, err := cache.Pop()
		if err != nil {
			return nil, err
		}

		if !ok {
			runtime.Gosched()
		} else {
			// saving the msg
			// has been check when put it into cache queue
			// so we don't need to check again
			inputMsg := msg.(*pipeline.Message)
			idx := inputMsg.GetSequence()
			fmt.Printf("[pipeline msg] getCompleteScopeDate - current idx %d, put it into msgVec[%d]\n", idx, idx)
			if msgVec[idx] == nil {
				msgVec[idx] = inputMsg
				cnt++
			} else {
				fmt.Printf("[pipeline msg] data compare: %d. uuid old = %s, new = %s\n", bytes.Compare(msgVec[idx].Data, inputMsg.Data), msgVec[idx].Uuid, inputMsg.Uuid)
				return nil, moerr.NewInternalErrorNoCtx("duplicate msg in morpc message cache")
			}
		}
	}
	fmt.Printf("[pipeline msg] getCompleteScopeDate - all seperate pipeine msg receive done. in cache msgCount = %d\n", msgCount)

	var dataBuffer []byte
	for i := range msgVec {
		dataBuffer = append(dataBuffer, msgVec[i].Data...)
	}
	dataBuffer = append(dataBuffer, m.Data...)
	return dataBuffer, nil
}

func (receiver *messageReceiverOnServer) acquireMessage() (*pipeline.Message, error) {
	message, ok := receiver.messageAcquirer().(*pipeline.Message)
	if !ok {
		return nil, moerr.NewInternalError(receiver.ctx, "get a message with wrong type.")
	}
	message.SetID(receiver.messageId)
	return message, nil
}

// newCompile make and return a new compile to run a pipeline.
func (receiver *messageReceiverOnServer) newCompile() *Compile {
	// compile is almost surely wanting a small or middle pool.  Later.
	mp, err := mpool.NewMPool("compile", 0, mpool.NoFixed)
	if err != nil {
		panic(err)
	}
	pHelper, cnInfo := receiver.procBuildHelper, receiver.cnInformation
	proc := process.New(
		receiver.ctx,
		mp,
		pHelper.txnClient,
		pHelper.txnOperator,
		cnInfo.fileService)
	proc.UnixTime = pHelper.unixTime
	proc.Id = pHelper.id
	proc.Lim = pHelper.lim
	proc.SessionInfo = pHelper.sessionInfo
	proc.AnalInfos = make([]*process.AnalyzeInfo, len(pHelper.analysisNodeList))
	for i := range proc.AnalInfos {
		proc.AnalInfos[i].NodeId = pHelper.analysisNodeList[i]
	}
	proc.DispatchNotifyCh = make(chan process.WrapCs, 1)

	c := &Compile{
		ctx:  receiver.ctx,
		proc: proc,
		e:    cnInfo.storeEngine,
		anal: &anaylze{},
		addr: colexec.CnAddr,
	}

	c.fill = func(_ any, b *batch.Batch) error {
		return receiver.sendBatch(b)
	}
	return c
}

func (receiver *messageReceiverOnServer) sendError(
	errInfo error) error {
	message, err := receiver.acquireMessage()
	if err != nil {
		return err
	}
	message.SetID(receiver.messageId)
	message.SetSid(pipeline.MessageEnd)
	if errInfo != nil {
		message.SetMoError(receiver.ctx, errInfo)
	}
	return receiver.clientSession.Write(receiver.ctx, message)
}

func (receiver *messageReceiverOnServer) sendBatch(
	b *batch.Batch) error {
	// there's no need to send the nil batch.
	if b == nil {
		return nil
	}
	data, err := types.Encode(b)
	if err != nil {
		return err
	}

	checksum := crc32.ChecksumIEEE(data)
	if len(data) <= receiver.maxMessageSize {
		m, errA := receiver.acquireMessage()
		if errA != nil {
			return errA
		}
		m.SetData(data)
		// XXX too bad.
		m.SetMessageType(pipeline.BatchMessage)
		m.SetCheckSum(checksum)
		m.SetSequence(receiver.sequence)
		receiver.sequence++
		return receiver.clientSession.Write(receiver.ctx, m)
	}
	// if data is too large, cut and send
	for start, end := 0, 0; start < len(data); start = end {
		m, errA := receiver.acquireMessage()
		if errA != nil {
			return errA
		}
		end = start + receiver.maxMessageSize
		if end > len(data) {
			end = len(data)
			m.SetSid(pipeline.Last)
		} else {
			m.SetSid(pipeline.WaitingNext)
		}
		m.SetMessageType(pipeline.BatchMessage)
		m.SetData(data[start:end])
		m.SetCheckSum(checksum)
		m.SetSequence(receiver.sequence)
		receiver.sequence++

		if errW := receiver.clientSession.Write(receiver.ctx, m); errW != nil {
			return errW
		}
	}
	return nil
}

func (receiver *messageReceiverOnServer) sendEndMessage() error {
	message, err := receiver.acquireMessage()
	if err != nil {
		return err
	}
	message.SetSid(pipeline.MessageEnd)
	message.SetID(receiver.messageId)
	message.SetMessageType(receiver.messageTyp)

	analysisInfo := receiver.finalAnalysisInfo
	if len(analysisInfo) > 0 {
		anas := &pipeline.AnalysisList{
			List: make([]*plan.AnalyzeInfo, len(analysisInfo)),
		}
		for i, a := range analysisInfo {
			anas.List[i] = convertToPlanAnalyzeInfo(a)
		}
		data, err := anas.Marshal()
		if err != nil {
			return err
		}
		message.SetAnalysis(data)
	}
	return receiver.clientSession.Write(receiver.ctx, message)
}

func generateProcessHelper(data []byte, cli client.TxnClient) (processHelper, error) {
	procInfo := &pipeline.ProcessInfo{}
	err := procInfo.Unmarshal(data)
	if err != nil {
		return processHelper{}, err
	}

	result := processHelper{
		id:               procInfo.Id,
		lim:              convertToProcessLimitation(procInfo.Lim),
		unixTime:         procInfo.UnixTime,
		txnClient:        cli,
		analysisNodeList: procInfo.GetAnalysisNodeList(),
	}
	result.txnOperator, err = cli.NewWithSnapshot([]byte(procInfo.Snapshot))
	if err != nil {
		return processHelper{}, err
	}
	result.sessionInfo, err = convertToProcessSessionInfo(procInfo.SessionInfo)
	if err != nil {
		return processHelper{}, err
	}

	return result, nil
}
