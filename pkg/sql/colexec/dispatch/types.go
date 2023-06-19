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

package dispatch

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/common/morpc"
	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const (
	maxMessageSizeToMoRpc = 64 * mpool.MB
	procTimeout           = 10000 * time.Second
	waitNotifyTimeout     = 40 * time.Second
	shuffleBatchSize      = 1024 * 8 //8k

	// send to all reg functions
	SendToAllLocalFunc = iota
	SendToAllRemoteFunc
	SendToAllFunc

	// send to any reg functions
	SendToAnyLocalFunc
	SendToAnyRemoteFunc
	SendToAnyFunc

	//shuffle to all reg functions
	ShuffleToAllFunc
)

type WrapperClientSession struct {
	msgId uint64
	cs    morpc.ClientSession
	uuid  uuid.UUID
	// toAddr string
	doneCh chan bool
}
type container struct {
	// the clientsession info for the channel you want to dispatch
	remoteReceivers []*WrapperClientSession
	// sendFunc is the rule you want to send batch
	sendFunc func(bat *batch.Batch, ap *Argument, proc *process.Process) (bool, error)

	// isRemote specify it is a remote receiver or not
	isRemote bool
	// prepared specify waiting remote receiver ready or not
	prepared bool
	end      bool

	// isParallel specify it is parallel or not
	// if yes, close with the parallel-close way
	isParallel bool

	// for send-to-any function decide send to which reg
	sendCnt       int
	aliveRegCnt   int
	localRegsCnt  int
	remoteRegsCnt int

	// for shuffle reuse memory
	sels         [][]int32
	remoteToIdx  map[uuid.UUID]int
	shuffledBats []*batch.Batch
	batsCount    int
}

type Argument struct {
	ctr *container

	// IsSink means this is a Sink Node
	IsSink bool
	// FuncId means the sendFunc you want to call
	FuncId int
	// LocalRegs means the local register you need to send to.
	LocalRegs []*process.WaitRegister
	// RemoteRegs specific the remote reg you need to send to.
	RemoteRegs []colexec.ReceiveInfo
	// for shuffle
	ShuffleColIdx       int32
	ShuffleType         int32
	ShuffleColMin       int64
	ShuffleColMax       int64
	ShuffleRegIdxLocal  []int
	ShuffleRegIdxRemote []int

	// for parallel dispatch
	ParallelNum *int32
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	for _, r := range arg.ctr.remoteReceivers {
		r.doneCh <- pipelineFailed
	}

	if arg.ctr.isParallel {
		if atomic.AddInt32(arg.ParallelNum, -1) > 0 {
			fmt.Printf("[dispatch.Free] (%t) proc %p free parallel type (not the last one)\n", pipelineFailed, proc)
			return
		}
		fmt.Printf("[dispatch.Free] (%t) proc %p free parallel type (last one)\n", pipelineFailed, proc)
	}

	/*
		if arg.ctr.isRemote {
			if !arg.ctr.prepared {
				arg.waitRemoteRegsReady(proc)
			}
			for _, r := range arg.ctr.remoteReceivers {
				timeoutCtx, cancel := context.WithTimeout(context.Background(), procTimeout)
				_ = cancel
				message := cnclient.AcquireMessage()
				{
					message.Id = r.msgId
					message.Cmd = pipeline.BatchMessage
					message.Sid = pipeline.MessageEnd
					message.Uuid = r.uuid[:]
				}
				if pipelineFailed {
					err := moerr.NewInternalError(proc.Ctx, "pipeline failed")
					message.Err = pipeline.EncodedMessageError(timeoutCtx, err)
				}
				r.cs.Write(timeoutCtx, message)
				//colexec.Srv.CleanUuidFromMap(r.uuid)
				close(r.doneCh)
			}
		}
	*/

	for i := range arg.LocalRegs {
		if !pipelineFailed {
			select {
			case <-arg.LocalRegs[i].Ctx.Done():
			case arg.LocalRegs[i].Ch <- nil:
			}
		}
		close(arg.LocalRegs[i].Ch)
	}
}
