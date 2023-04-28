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

package connector

import (
	"bytes"

	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(arg any, buf *bytes.Buffer) {
	buf.WriteString("pipe connector")
}

func Prepare(_ *process.Process, _ any) error {
	return nil
}

func Call(_ int, proc *process.Process, arg any, _ bool, _ bool) (bool, error) {
	ap := arg.(*Argument)
	bat := proc.InputBatch()
	if bat == nil {
		return true, nil
	}
	if bat.Length() == 0 {
		return false, nil
	}

	// do not send the source batch to remote node.
	for i := range bat.Vecs {
		if bat.Vecs[i].NeedDup() {
			oldVec := bat.Vecs[i]
			vec, err := oldVec.Dup(proc.Mp())
			if err != nil {
				return false, err
			}
			bat.ReplaceVector(oldVec, vec)
			oldVec.Free(proc.Mp())
		}
	}

	return ap.Send(bat, proc)
}

func (arg *Argument) Send(bat *batch.Batch, proc *process.Process) (bool, error) {
	reg := arg.Reg
	select {
	case <-reg.Ctx.Done():
		proc.PutBatch(bat)
		return true, nil
	case reg.Ch <- bat:
		proc.SetInputBatch(nil)
		return false, nil
	}
}
