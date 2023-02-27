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

package offset

import (
	"bytes"
	"fmt"

	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

var emptyBatch = &batch.Batch{}

func String(arg any, buf *bytes.Buffer) {
	ap := arg.(*Argument)
	buf.WriteString(fmt.Sprintf("offset(%v)", ap.Offset))
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.bat = batch.NewWithSize(len(ap.Types))
	ap.ctr.vecs = make([]*vector.Vector, len(ap.Types))
	for i := range ap.Types {
		vec := vector.New(ap.Types[i])
		ap.ctr.vecs[i] = vec
		ap.ctr.bat.SetVector(int32(i), vec)
	}
	// init ufs
	return nil
}

func Call(idx int, proc *process.Process, arg any, isFirst bool, isLast bool) (bool, error) {
	bat := proc.InputBatch()
	if bat == nil {
		return true, nil
	}
	if bat.Length() == 0 {
		return false, nil
	}
	ap := arg.(*Argument)
	anal := proc.GetAnalyze(idx)
	anal.Start()
	defer anal.Stop()
	anal.Input(bat, isFirst)

	length := bat.Length()
	if ap.Seen > ap.Offset {
		ap.ctr.bat.Reset()
		for i, vec := range ap.ctr.vecs {
			uf := ap.ctr.ufs[i]
			srcVec := bat.GetVector(int32(i))
			for j := 0; i < length; j++ {
				if err := uf(vec, srcVec, int64(i)); err != nil {
					return false, err
				}
			}
		}
		ap.ctr.bat.Zs = append(ap.ctr.bat.Zs, bat.Zs...)
		return false, nil
	}
	if ap.Seen+uint64(length) > ap.Offset {
		ap.ctr.bat.Reset()
		start, count := int64(ap.Offset-ap.Seen), int64(length)-int64(ap.Offset-ap.Seen)
		for i, vec := range ap.ctr.vecs {
			uf := ap.ctr.ufs[i]
			srcVec := bat.GetVector(int32(i))
			for j := int64(0); j < count; j++ {
				if err := uf(vec, srcVec, j+start); err != nil {
					return false, err
				}
			}
		}
		for i := int64(0); i < count; i++ {
			ap.ctr.bat.Zs = append(ap.ctr.bat.Zs, i+start)
		}
		return false, nil
	}
	ap.Seen += uint64(length)
	proc.SetInputBatch(emptyBatch)
	return false, nil
}
