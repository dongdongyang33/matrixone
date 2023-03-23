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

package loopmark

import (
	"bytes"
	"time"

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(_ any, buf *bytes.Buffer) {
	buf.WriteString(" loop mark join ")
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.InitByTypes(ap.Types, proc)
	return nil
}

func Call(idx int, proc *process.Process, arg any, isFirst bool, isLast bool) (bool, error) {
	anal := proc.GetAnalyze(idx)
	anal.Start()
	defer anal.Stop()
	ap := arg.(*Argument)
	ctr := ap.ctr
	for {
		switch ctr.state {
		case Build:
			if err := ctr.build(ap, proc, anal); err != nil {
				ctr.state = End
				return false, err
			}
			ctr.state = Probe
		case Probe:
			start := time.Now()
			bat := <-proc.Reg.MergeReceivers[0].Ch
			anal.WaitStop(start)

			if bat == nil {
				ctr.state = End
				continue
			}
			if bat.Length() == 0 {
				bat.SubCnt(1)
				continue
			}

			err := ctr.probeFunc(bat, ap, proc, anal, isFirst, isLast)
			bat.SubCnt(1)
			return false, err
		default:
			proc.SetInputBatch(nil)
			return true, nil
		}
	}
}

func (ctr *container) build(ap *Argument, proc *process.Process, anal process.Analyze) error {
	start := time.Now()
	bat := <-proc.Reg.MergeReceivers[1].Ch
	anal.WaitStop(start)
	if bat != nil {
		ctr.bat = bat
		ctr.probeFunc = ctr.probe
	} else {
		ctr.probeFunc = ctr.emptyProbe
	}
	return nil
}

func (ctr *container) emptyProbe(bat *batch.Batch, ap *Argument, proc *process.Process, anal process.Analyze, isFirst bool, isLast bool) error {
	anal.Input(bat, isFirst)
	ctr.OutBat.Reset()
	for i, rp := range ap.Result {
		uf := ctr.Ufs[i]
		vec := ctr.OutBat.GetVector(int32(i))
		if rp >= 0 {
			srcVec := bat.GetVector(rp)
			for j := 0; j < bat.Length(); j++ {
				if err := uf(vec, srcVec, int64(j)); err != nil {
					return err
				}
			}
		} else {
			for j := 0; j < bat.Length(); j++ {
				if err := vector.AppendFixed(vec, false, false, proc.Mp()); err != nil {
					return err
				}
			}
		}
	}
	ctr.OutBat.Zs = append(ctr.OutBat.Zs, bat.Zs...)
	anal.Output(ctr.OutBat, isLast)
	proc.SetInputBatch(ctr.OutBat)
	return nil
}

func (ctr *container) probe(bat *batch.Batch, ap *Argument, proc *process.Process, anal process.Analyze, isFirst bool, isLast bool) error {
	anal.Input(bat, isFirst)
	markPos := -1
	for i, pos := range ap.Result {
		if pos == -1 {
			markPos = i
			break
		}
	}
	if markPos == -1 {
		return moerr.NewInternalError(proc.Ctx, "MARK join must output mark column")
	}
	count := bat.Length()
	for i := 0; i < count; i++ {
		vec, err := colexec.JoinFilterEvalExpr(bat, ctr.bat, i, proc, ap.Cond)
		if err != nil {
			return err
		}
		needFree := true
		for j := range bat.Vecs {
			if bat.Vecs[j] == vec {
				needFree = false
			}
		}
		exprVals := vector.MustFixedCol[bool](vec)
		hasTrue := false
		hasNull := false
		if vec.IsConst() {
			if vec.IsConstNull() {
				hasNull = true
			} else {
				hasTrue = exprVals[0]
			}
		} else {
			for j := range exprVals {
				if vec.GetNulls().Contains(uint64(j)) {
					hasNull = true
				} else if exprVals[j] {
					hasTrue = true
				}
			}
		}
		if hasTrue {
			vector.AppendFixed(ctr.OutBat.Vecs[markPos], true, false, proc.Mp())
		} else if hasNull {
			vector.AppendFixed(ctr.OutBat.Vecs[markPos], false, true, proc.Mp())
		} else {
			vector.AppendFixed(ctr.OutBat.Vecs[markPos], false, false, proc.Mp())
		}
		if needFree {
			vec.Free(proc.Mp())
		}
	}
	for i, rp := range ap.Result {
		uf := ctr.Ufs[i]
		vec := ctr.OutBat.GetVector(int32(i))
		if rp >= 0 {
			srcVec := bat.GetVector(rp)
			for j := 0; j < bat.Length(); j++ {
				if err := uf(vec, srcVec, int64(j)); err != nil {
					return err
				}
			}
		}
	}
	ctr.OutBat.Zs = append(ctr.OutBat.Zs, bat.Zs...)
	anal.Output(ctr.OutBat, isLast)
	proc.SetInputBatch(ctr.OutBat)
	return nil
}
