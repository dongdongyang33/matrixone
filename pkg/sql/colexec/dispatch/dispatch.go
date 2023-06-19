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
	"bytes"
	"context"
	"fmt"
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("dispatch")
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ctr := new(container)
	ap.ctr = ctr
	ctr.localRegsCnt = len(ap.LocalRegs)
	ctr.remoteRegsCnt = len(ap.RemoteRegs)
	ctr.aliveRegCnt = ctr.localRegsCnt + ctr.remoteRegsCnt
	if ap.ParallelNum != nil {
		fmt.Printf("[dispatch.prepared] proc %p with parallel prepared\n", proc)
		ctr.isParallel = atomic.LoadInt32(ap.ParallelNum) > 0
	}

	switch ap.FuncId {
	case SendToAllFunc:
		if ctr.remoteRegsCnt == 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAllFunc should include RemoteRegs")
		}
		if len(ap.LocalRegs) == 0 {
			ctr.sendFunc = sendToAllRemoteFunc
		} else {
			ctr.sendFunc = sendToAllFunc
		}
		if err := ap.prepareRemote(proc, false); err != nil {
			return err
		}

	case ShuffleToAllFunc:
		ap.ctr.sendFunc = shuffleToAllFunc
		if ap.ctr.remoteRegsCnt > 0 {
			if err := ap.prepareRemote(proc, false); err != nil {
				return err
			}
		} else {
			ap.prepareLocal()
		}
		ap.initShuffle()

	case SendToAnyFunc:
		if ctr.remoteRegsCnt == 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAnyFunc should include RemoteRegs")
		}
		if len(ap.LocalRegs) == 0 {
			ctr.sendFunc = sendToAnyRemoteFunc
		} else {
			ctr.sendFunc = sendToAnyFunc
		}
		if err := ap.prepareRemote(proc, true); err != nil {
			return err
		}

		// Update the remote count because sendToAny allows some receiver failed
		ctr.remoteRegsCnt = len(ap.ctr.remoteReceivers)
		ap.ctr.aliveRegCnt = ap.ctr.remoteRegsCnt + ap.ctr.localRegsCnt

	case SendToAllLocalFunc:
		if ctr.remoteRegsCnt != 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAllLocalFunc should not send to remote")
		}
		ctr.sendFunc = sendToAllLocalFunc
		ap.prepareLocal()

	case SendToAnyLocalFunc:
		if ctr.remoteRegsCnt != 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAnyLocalFunc should not send to remote")
		}
		ap.ctr.sendFunc = sendToAnyLocalFunc
		ap.prepareLocal()

	default:
		return moerr.NewInternalError(proc.Ctx, "wrong sendFunc id for dispatch")
	}

	return nil
}

func Call(idx int, proc *process.Process, arg any, isFirst bool, isLast bool) (bool, error) {
	ap := arg.(*Argument)

	bat := proc.InputBatch()
	if bat == nil || ap.ctr.end {
		if ap.FuncId == ShuffleToAllFunc {
			return sendShuffledBats(ap, proc)
		}
		return true, nil
	}

	if bat.Length() == 0 {
		bat.Clean(proc.Mp())
		return false, nil
	}
	return ap.ctr.sendFunc(bat, ap, proc)
}

func (arg *Argument) waitRemoteRegsReady(proc *process.Process, isAny bool) error {
	cnt := len(arg.RemoteRegs)
	fmt.Printf("[dispatch.Wait] proc %p waiting ...\n", proc)
	for cnt > 0 {
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), waitNotifyTimeout)
		defer timeoutCancel()
		select {
		case <-timeoutCtx.Done():
			if !isAny {
				fmt.Printf("[dispatch.Wait] proc %p waiting timeout\n", proc)
				logutil.Errorf("waiting notify msg timeout")
				return moerr.NewInternalErrorNoCtx("wait notify message timeout")
			}
		case <-proc.Ctx.Done():
			arg.ctr.end = true
			fmt.Printf("[dispatch.Wait] proc %p ctx done\n", proc)
			logutil.Warn("conctx done during dispatch")
			break
		case csinfo := <-proc.DispatchNotifyCh:
			arg.ctr.remoteReceivers = append(arg.ctr.remoteReceivers, &WrapperClientSession{
				msgId:  csinfo.MsgId,
				cs:     csinfo.Cs,
				uuid:   csinfo.Uid,
				doneCh: csinfo.DoneCh,
			})
		}
		cnt--
	}
	return nil
}

func (arg *Argument) prepareRemote(proc *process.Process, isAny bool) error {
	arg.ctr.isRemote = true
	arg.ctr.remoteReceivers = make([]*WrapperClientSession, 0, arg.ctr.remoteRegsCnt)
	arg.ctr.remoteToIdx = make(map[uuid.UUID]int)
	for i, rr := range arg.RemoteRegs {
		if arg.FuncId == ShuffleToAllFunc {
			arg.ctr.remoteToIdx[rr.Uuid] = arg.ShuffleRegIdxRemote[i]
		}
		if !arg.ctr.isParallel {
			colexec.Srv.PutParallelNumForUuid(rr.Uuid, 1)
		} else {
			colexec.Srv.PutParallelNumForUuid(rr.Uuid, int(*arg.ParallelNum))
		}
		colexec.Srv.PutProcIntoUuidMap(rr.Uuid, proc)
	}

	return arg.waitRemoteRegsReady(proc, isAny)
}

func (arg *Argument) prepareLocal() {
	arg.ctr.prepared = true
	arg.ctr.isRemote = false
	arg.ctr.remoteReceivers = nil
}

func (arg *Argument) initShuffle() {
	if arg.ctr.sels == nil {
		arg.ctr.sels = make([][]int32, arg.ctr.aliveRegCnt)
		for i := 0; i < arg.ctr.aliveRegCnt; i++ {
			arg.ctr.sels[i] = make([]int32, 8192)
		}
		arg.ctr.batsCount = 0
		arg.ctr.shuffledBats = make([]*batch.Batch, arg.ctr.aliveRegCnt)
	}
}

func (arg *Argument) getSels() [][]int32 {
	for i := range arg.ctr.sels {
		arg.ctr.sels[i] = arg.ctr.sels[i][:0]
	}
	return arg.ctr.sels
}
