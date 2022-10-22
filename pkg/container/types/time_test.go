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

package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTime_StringAndString2(t *testing.T) {
	testCases := []struct {
		name       string
		input      Time
		strExpect  string
		str2Expect string
		precision  int32
	}{
		{
			name: "TestString-NoPrecision",
			// 11:22:33
			input:      FromTimeClock(false, 11, 22, 33, 0),
			strExpect:  "11:22:33",
			str2Expect: "11:22:33",
			precision:  0,
		},
		{
			name: "TestString-Precision",
			// 11:22:33.123
			input:      FromTimeClock(false, 11, 22, 33, 123000),
			strExpect:  "11:22:33",
			str2Expect: "11:22:33.12300",
			precision:  5,
		},
		{
			name: "TestString-ShortterPrecision",
			// 11:22:33.123
			input:      FromTimeClock(false, 11, 22, 33, 123000),
			strExpect:  "11:22:33",
			str2Expect: "11:22:33.12",
			precision:  2,
		},
		{
			name: "TestString-Minus",
			// 11:22:33.125000
			input:      FromTimeClock(true, 11, 22, 33, 125000),
			strExpect:  "-11:22:33",
			str2Expect: "-11:22:33.12",
			precision:  2,
		},
	}

	for _, v := range testCases {
		// only 1 input
		strActual := v.input.String()
		require.Equal(t, strActual, v.strExpect)
		str2Actual := v.input.String2(v.precision)
		require.Equal(t, str2Actual, v.str2Expect)
	}
}

func TestTime_ParseTime(t *testing.T) {
	testCases := []struct {
		name      string
		inputStr  string
		expected  Time
		precision int32
	}{
		{
			name: "TestString-NoPrecision",
			// 11:22:33
			expected:  FromTimeClock(false, 11, 22, 33, 0),
			precision: 0,
		},
		{
			name: "TestString-Precision",
			// 11:22:33.123
			expected:  FromTimeClock(false, 11, 22, 33, 123000),
			precision: 5,
		},
		{
			name: "TestString-ShortterPrecision",
			// 11:22:33.123
			expected:  FromTimeClock(false, 11, 22, 33, 123000),
			precision: 2,
		},
		{
			name: "TestString-Minus",
			// 11:22:33.125000
			expected:  FromTimeClock(true, 11, 22, 33, 125000),
			precision: 2,
		},
	}

	for _, v := range testCases {
		// only 1 input
		strActual := v.expected.String()
		require.Equal(t, strActual, v.strExpect)
		str2Actual := v.expected.String2(v.precision)
		require.Equal(t, str2Actual, v.str2Expect)
	}
}
