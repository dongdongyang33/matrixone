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

	"github.com/matrixorigin/matrixone/pkg/common/moerr"
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
		isErr     bool
	}{
		// ==================== Date format: yyyy-mm-dd hh:mm:ss(.msec) ====================
		{
			name: "TestParse-NoPrecision",
			// 11:22:33
			inputStr:  "2022-12-12 11:22:33",
			expected:  FromTimeClock(false, 11, 22, 33, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse-Precision01",
			// 11:22:33
			inputStr:  "2022-12-12 11:22:33.1234",
			expected:  FromTimeClock(false, 11, 22, 33, 123000),
			precision: 3,
			isErr:     false,
		},
		{
			name: "TestParse-Precision02",
			// 11:22:33
			inputStr:  "2022-12-12 11:22:33.1235",
			expected:  FromTimeClock(false, 11, 22, 33, 124000),
			precision: 3,
			isErr:     false,
		},
		{
			name: "TestParse-DateError",
			// 11:22:33
			inputStr: "2022-12-33 11:22:33",
			isErr:    true,
		},
		// ==================== Date format: yyyymmddhhmmss(.msec) ====================
		{
			name: "TestParse2-NoPrecision",
			// 11:22:33
			inputStr:  "20221212112233",
			expected:  FromTimeClock(false, 11, 22, 33, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse2-Precision01",
			// 11:22:33
			inputStr:  "20221212112233.1234",
			expected:  FromTimeClock(false, 11, 22, 33, 123000),
			precision: 3,
			isErr:     false,
		},
		{
			name: "TestParse2-Precision02",
			// 11:22:33
			inputStr:  "20221212112233.1235",
			expected:  FromTimeClock(false, 11, 22, 33, 124000),
			precision: 3,
			isErr:     false,
		},
		{
			name: "TestParse2-DateError",
			// 11:22:33
			inputStr: "20221233112233",
			isErr:    true,
		},
		// ==================== Time format: hh:mm:ss(.msec) ====================
		{
			name: "TestParse3-NoPrecision",
			// 11:22:33
			inputStr:  "11:22:33",
			expected:  FromTimeClock(false, 11, 22, 33, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse3-NoPrecision2",
			// 11:22:33
			inputStr:  "555:22:33",
			expected:  FromTimeClock(false, 555, 22, 33, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse3-NoPrecision2",
			// 11:22:33
			inputStr:  "-555:22:33",
			expected:  FromTimeClock(true, 555, 22, 33, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse3-Precision",
			// 11:22:33
			inputStr:  "11:22:33.1234",
			expected:  FromTimeClock(false, 11, 22, 33, 123000),
			precision: 3,
			isErr:     false,
		},
		{
			name: "TestParse3-Precision",
			// 11:22:33
			inputStr:  "11:22:33.1235",
			expected:  FromTimeClock(false, 11, 22, 33, 124000),
			precision: 3,
			isErr:     false,
		},
		// ==================== Time format: hhmmss(.msec) ====================
		{
			name: "TestParse4-NoPrecision",
			// 11:22:33
			inputStr:  "112",
			expected:  FromTimeClock(false, 0, 1, 12, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse4-NoPrecision",
			// 11:22:33
			inputStr:  "-112",
			expected:  FromTimeClock(true, 0, 1, 12, 0),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse4-NoPrecision",
			// 11:22:33
			inputStr:  "-11232",
			expected:  FromTimeClock(true, 0, 1, 12, 32),
			precision: 0,
			isErr:     false,
		},
		{
			name: "TestParse4-Precision",
			// 11:22:33
			inputStr:  "-11232",
			expected:  FromTimeClock(true, 0, 1, 12, 32),
			precision: 0,
			isErr:     false,
		},
	}

	for _, v := range testCases {
		// only 1 input
		parsed, err := ParseTime(v.inputStr, v.precision)
		if !v.isErr {
			require.NoError(t, err)
			require.Equal(t, parsed, v.expected)
		} else {
			require.Equal(t, err, moerr.NewInvalidInput("invalid time value %s", v.inputStr))
		}
	}
}
