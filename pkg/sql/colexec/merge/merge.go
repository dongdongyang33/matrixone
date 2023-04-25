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
	"github.com/matrixorigin/matrixone/pkg/logutil"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func String(_ any, buf *bytes.Buffer) {
	buf.WriteString(" union all ")
}

func Prepare(proc *process.Process, arg any) error {
	ap := arg.(*Argument)
	ap.ctr = new(container)
	ap.ctr.receiverListener = make([]reflect.SelectCase, len(proc.Reg.MergeReceivers))
	ap.ctr.test = make([]int, len(proc.Reg.MergeReceivers))
	str := ""
	for i, mr := range proc.Reg.MergeReceivers {
		if i != 0 {
			str += " ,"
		}
		str += fmt.Sprintf("%p", proc.Reg.MergeReceivers[i].Ch)
		ap.ctr.receiverListener[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(mr.Ch),
		}
		ap.ctr.test[i] = i
	}
	fmt.Printf("[merge.Prepare] proc = %p, with chs = %s\n", proc, str)

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
		select {
		case <-proc.Ctx.Done():
			moerr.NewInternalErrorNoCtx("merge err...")
		default:
		}

		start := time.Now()
		chosen, value, ok := reflect.Select(ctr.receiverListener)
		if !ok {
			fmt.Printf("[merge.Call] proc %p close ch %p\n", proc, proc.Reg.MergeReceivers[ap.ctr.test[chosen]].Ch)
			ctr.receiverListener = append(ctr.receiverListener[:chosen], ctr.receiverListener[chosen+1:]...)
			ap.ctr.test = append(ctr.test[:chosen], ctr.test[chosen+1:]...)
			logutil.Errorf("pipeline closed unexpectedly")
			return true, nil
		}
		anal.WaitStop(start)

		pointer := value.UnsafePointer()
		bat := (*batch.Batch)(pointer)
		if bat == nil {
			fmt.Printf("[merge.Call] proc %p receive nil from ch %p\n", proc, proc.Reg.MergeReceivers[ap.ctr.test[chosen]].Ch)
			ap.ctr.test = append(ctr.test[:chosen], ctr.test[chosen+1:]...)
			ctr.receiverListener = append(ctr.receiverListener[:chosen], ctr.receiverListener[chosen+1:]...)
			ctr.aliveMergeReceiver--
			continue
		}
		if bat.Length() == 0 {
			bat.Clean(proc.Mp())
			continue
		}
		anal.Input(bat, isFirst)
		anal.Output(bat, isLast)
		proc.SetInputBatch(bat)
		return false, nil
	}
}
