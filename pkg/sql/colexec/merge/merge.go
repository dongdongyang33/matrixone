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

package merge

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(_ any, buf *bytes.Buffer) {
	buf.WriteString(" union all ")
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.receiverListener = make([]reflect.SelectCase, len(proc.Reg.MergeReceivers))
	for i, mr := range proc.Reg.MergeReceivers {
		ap.ctr.receiverListener[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(mr.Ch),
		}
		ap.ctr.debug = append(ap.ctr.debug, i)
	}

	ap.ctr.aliveMergeReceiver = len(proc.Reg.MergeReceivers)
	return nil
}

func Call(idx int, proc *process.Process, arg any, isFirst bool, isLast bool) (bool, error) {
	anal := proc.GetAnalyze(idx)
	anal.Start()
	defer anal.Stop()
	ap := arg.(*Argument)
	ctr := ap.ctr

	for {
		if ctr.aliveMergeReceiver == 0 {
			proc.SetInputBatch(nil)
			return true, nil
		}

		fmt.Printf("[mergemerge] current receiver %s. waiting ... proc = %p\n", ctr.getCurrentString(), proc)
		start := time.Now()
		chosen, value, ok := reflect.Select(ctr.receiverListener)
		if !ok {
			return false, moerr.NewInternalError(proc.Ctx, "pipeline closed unexpectedly")
		}
		anal.WaitStop(start)

		pointer := value.UnsafePointer()
		bat := (*batch.Batch)(pointer)
		fmt.Printf("[mergemerge] receiver %s receive a msg. proc = %p\n", ctr.getCurrentString(), proc)
		if bat == nil {
			ctr.receiverListener = append(ctr.receiverListener[:chosen], ctr.receiverListener[chosen+1:]...)
			ctr.aliveMergeReceiver--

			ctr.debug = append(ctr.debug[:chosen], ctr.debug[chosen+1:]...)
			fmt.Printf("[mergemerge] %d closed. current receiver %s. proc = %p\n", chosen, ctr.getCurrentString(), proc)
			continue
		}
		if bat.Length() == 0 {
			continue
		}
		anal.Input(bat, isFirst)
		anal.Output(bat, isLast)
		proc.SetInputBatch(bat)
		return false, nil
	}
}

func (ctr *container) getCurrentString() string {
	s := "["
	for i, n := range ctr.debug {
		if i != 0 {
			s += " ,"
		}
		s += fmt.Sprintf("%d", n)
	}
	s += "]"
	return s
}
