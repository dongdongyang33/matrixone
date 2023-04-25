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
	"fmt"
	"reflect"

	"github.com/matrixorigin/matrixone/pkg/container/batch"

	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

type container struct {
	aliveMergeReceiver int
	// receiverListener is a structure to listen all the merge receiver.
	receiverListener []reflect.SelectCase
	test             []int
}

type Argument struct {
	ctr *container
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	if arg.ctr != nil {
		listeners := arg.ctr.receiverListener
		alive := len(listeners)
		fmt.Printf("[merge.Free] proc %p alive = %d, failed = %t\n", proc, alive, pipelineFailed)
		for alive != 0 {
			chosen, value, ok := reflect.Select(listeners)
			if !ok {
				fmt.Printf("[merge.Free] proc %p close ch %p\n", proc, proc.Reg.MergeReceivers[arg.ctr.test[chosen]].Ch)
				listeners = append(listeners[:chosen], listeners[chosen+1:]...)
				arg.ctr.test = append(arg.ctr.test[:chosen], arg.ctr.test[chosen+1:]...)
				alive--
				continue
			}
			pointer := value.UnsafePointer()
			bat := (*batch.Batch)(pointer)
			if bat == nil {
				fmt.Printf("[merge.Free] proc %p receive nil from ch %p\n", proc, proc.Reg.MergeReceivers[arg.ctr.test[chosen]].Ch)
				alive--
				listeners = append(listeners[:chosen], listeners[chosen+1:]...)
				arg.ctr.test = append(arg.ctr.test[:chosen], arg.ctr.test[chosen+1:]...)
				continue
			}
			bat.Clean(proc.Mp())
		}
	}
}
