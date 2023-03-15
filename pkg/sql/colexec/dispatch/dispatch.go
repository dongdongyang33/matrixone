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

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/defines"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("dispatch")
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.bat = batch.NewWithSize(len(ap.Types))
	ap.ctr.vecs = make([]*vector.Vector, len(ap.Types))
	for i := range ap.Types {
		vec := vector.NewVec(ap.Types[i])
		vec.PreExtend(defines.DefaultVectorRows, proc.Mp())
		ap.ctr.vecs[i] = vec
		ap.ctr.bat.SetVector(int32(i), vec)
	}
	ap.ctr.ufs = make([]func(*vector.Vector, *vector.Vector, int64) error, len(ap.Types))
	for i := range ap.Types {
		ap.ctr.ufs[i] = vector.GetUnionOneFunction(ap.Types[i], proc.Mp())
	}

	switch ap.FuncId {
	case SendToAllFunc:
		if len(ap.RemoteRegs) == 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAllFunc should include RemoteRegs")
		}
		ap.prepared = false
		ap.ctr.remoteReceivers = make([]*WrapperClientSession, 0, len(ap.RemoteRegs))
		ap.ctr.sendFunc = sendToAllFunc
		for _, rr := range ap.RemoteRegs {
			colexec.Srv.PutNotifyChIntoUuidMap(rr.Uuid, proc.DispatchNotifyCh)
		}

	case SendToAllLocalFunc:
		if len(ap.RemoteRegs) != 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAllLocalFunc should not send to remote")
		}
		ap.prepared = true
		ap.ctr.remoteReceivers = nil
		ap.ctr.sendFunc = sendToAllLocalFunc
	case SendToAnyLocalFunc:
		if len(ap.RemoteRegs) != 0 {
			return moerr.NewInternalError(proc.Ctx, "SendToAnyLocalFunc should not send to remote")
		}
		ap.prepared = true
		ap.ctr.remoteReceivers = nil
		ap.ctr.sendFunc = sendToAnyLocalFunc
	default:
		return moerr.NewInternalError(proc.Ctx, "wrong sendFunc id for dispatch")
	}

	return nil
}

func Call(idx int, proc *process.Process, arg any, isFirst bool, isLast bool) (bool, error) {
	ap := arg.(*Argument)

	// waiting all remote receive prepared
	// put it in Call() for better parallel

	bat := proc.InputBatch()
	if bat == nil {
		return true, nil
	}
	length := bat.Length()
	if length == 0 {
		return false, nil
	}

	ap.ctr.bat.Reset()
	for i, vec := range ap.ctr.vecs {
		uf := ap.ctr.ufs[i]
		srcVec := bat.GetVector(int32(i))
		for j := int64(0); j < int64(length); j++ {
			if err := uf(vec, srcVec, j); err != nil {
				return false, err
			}
		}
	}

	if err := ap.ctr.sendFunc(ap.ctr.bat, ap, proc); err != nil {
		return false, err
	}
	return false, nil
}

func (arg *Argument) waitRemoteRegsReady(proc *process.Process) {
	cnt := len(arg.RemoteRegs)
	for cnt > 0 {
		csinfo := <-proc.DispatchNotifyCh
		arg.ctr.remoteReceivers = append(arg.ctr.remoteReceivers, &WrapperClientSession{
			msgId:  csinfo.MsgId,
			cs:     csinfo.Cs,
			uuid:   csinfo.Uid,
			doneCh: csinfo.DoneCh,
		})
		cnt--
	}
	arg.prepared = true
}
