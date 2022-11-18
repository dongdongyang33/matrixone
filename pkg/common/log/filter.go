// Copyright 2022 Matrix Origin
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

package log

import (
	"sync/atomic"
)

var (
	// all log filters register here.
	filters = []logFilter{sampleFilter}

	// all SampleType register here.
	samples = map[SampleType]*sampleValue{
		ExampleSample: {},
	}
)

func sampleFilter(ctx LogContext) bool {
	if ctx.sampleType == noneSample {
		return true
	}
	if v, ok := samples[ctx.sampleType]; ok {
		return v.incr()%ctx.samplefrequency == 0
	}
	return false
}

type sampleValue struct {
	v uint64
}

func (s *sampleValue) incr() uint64 {
	return atomic.AddUint64(&s.v, 1)
}