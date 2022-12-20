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

package process

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/matrixorigin/matrixone/pkg/container/batch"
)

func NewAnalyzeInfo(nodeId int32) *AnalyzeInfo {
	return &AnalyzeInfo{
		NodeId:       nodeId,
		InputRows:    0,
		OutputRows:   0,
		TimeConsumed: 0,
		InputSize:    0,
		OutputSize:   0,
		MemorySize:   0,
	}
}

func (a *analyze) Start() {
	a.start = time.Now()
	if a.analInfo.Op != "unknow" {
		println(fmt.Sprintf("[%d] Op: %s, start time: %s", a.analInfo.formID, a.analInfo.Op, a.start.String()))
	} else {
		println(fmt.Sprintf("[%d] start time: %s", a.analInfo.formID, a.start.String()))
	}
}

func (a *analyze) Stop() {
	if a.analInfo != nil {
		span := int64(time.Since(a.start) / time.Microsecond)
		if span > 92233000000000 {
			println(fmt.Sprintf("[%d] Op %s span is too large! a.start = %s", a.analInfo.formID, a.analInfo.Op, a.start.String()))
		}
		atomic.AddInt64(&a.analInfo.TimeConsumed, span)
		if a.analInfo.Op != "unknow" {
			println(fmt.Sprintf("[%d] Op: %s, time += %d", a.analInfo.formID, a.analInfo.Op, span))
		} else {
			println(fmt.Sprintf("[%d] time += %d", a.analInfo.formID, span))
		}
	}
}

func (a *analyze) Alloc(size int64) {
	if a.analInfo != nil {
		atomic.AddInt64(&a.analInfo.MemorySize, size)
	}
}

func (a *analyze) Input(bat *batch.Batch) {
	if a.analInfo != nil && bat != nil {
		atomic.AddInt64(&a.analInfo.InputSize, int64(bat.Size()))
		atomic.AddInt64(&a.analInfo.InputRows, int64(bat.Length()))
		if a.analInfo.Op != "unknow" {
			println(fmt.Sprintf("[%d] Op: %s, input size += %d, rows += %d", a.analInfo.formID, a.analInfo.Op, bat.Size(), bat.Length()))
		} else {
			println(fmt.Sprintf("[%d] input size += %d, rows += %d", a.analInfo.formID, bat.Size(), bat.Length()))
		}
	}
}

func (a *analyze) ActualInput(bat *batch.Batch) {
	if a.analInfo != nil && bat != nil {
		atomic.AddInt64(&a.analInfo.ActualInputSize, int64(bat.Size()))
		atomic.AddInt64(&a.analInfo.ActualInputRows, int64(bat.Length()))
		if a.analInfo.Op != "unknow" {
			println(fmt.Sprintf("[%d] *actual* Op: %s, input size += %d, rows += %d", a.analInfo.formID, a.analInfo.Op, bat.Size(), bat.Length()))
		} else {
			println(fmt.Sprintf("[%d] *actual* input size += %d, rows += %d", a.analInfo.formID, bat.Size(), bat.Length()))
		}
	}
}

func (a *analyze) Output(bat *batch.Batch) {
	if a.analInfo != nil && bat != nil {
		atomic.AddInt64(&a.analInfo.OutputSize, int64(bat.Size()))
		atomic.AddInt64(&a.analInfo.OutputRows, int64(bat.Length()))
		if a.analInfo.Op != "unknow" {
			println(fmt.Sprintf("[%d] Op: %s, output size += %d, rows += %d", a.analInfo.formID, a.analInfo.Op, bat.Size(), bat.Length()))
		} else {
			println(fmt.Sprintf("[%d] output size += %d, rows += %d", a.analInfo.formID, bat.Size(), bat.Length()))
		}
	}
}

func (a *analyze) ActualOutput(bat *batch.Batch) {
	if a.analInfo != nil && bat != nil {
		atomic.AddInt64(&a.analInfo.ActualOutputSize, int64(bat.Size()))
		atomic.AddInt64(&a.analInfo.ActualOutputRows, int64(bat.Length()))
		if a.analInfo.Op != "unknow" {
			println(fmt.Sprintf("[%d] *actual* Op: %s, output size += %d, rows += %d", a.analInfo.formID, a.analInfo.Op, bat.Size(), bat.Length()))
		} else {
			println(fmt.Sprintf("[%d] *actual* output size += %d, rows += %d", a.analInfo.formID, bat.Size(), bat.Length()))
		}
	}
}

func (a *analyze) AddWaitingTime(start time.Time) {
	a.waitingTime += int64(time.Since(start) / time.Microsecond)
}
