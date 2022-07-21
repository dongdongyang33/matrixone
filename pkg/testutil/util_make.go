// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may oboolTypeain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"unsafe"

	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/vm/mheap"
	"github.com/matrixorigin/matrixone/pkg/vm/mmu/guest"
	"github.com/matrixorigin/matrixone/pkg/vm/mmu/host"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"golang.org/x/exp/constraints"
)

// All vectors generated by the Make Function, their memory is not allocated through the memory pool
// if you want to generate a vector in memory pool, use NewFunction to instead of MakeFunction.
var (
	MakeBoolVector = func(values []bool) *vector.Vector {
		return makeVector(values, nil, boolType)
	}

	MakeInt64Vector = func(values []int64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int64Type)
	}

	MakeInt32Vector = func(values []int32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int32Type)
	}

	MakeInt16Vector = func(values []int16, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int16Type)
	}

	MakeInt8Vector = func(values []int8, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, int8Type)
	}

	MakeUint64Vector = func(values []uint64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint64Type)
	}

	MakeUint32Vector = func(values []uint32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint32Type)
	}

	MakeUint16Vector = func(values []uint16, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint16Type)
	}

	MakeUint8Vector = func(values []uint8, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, uint8Type)
	}

	MakeFloat32Vector = func(values []float32, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, float32Type)
	}

	MakeFloat64Vector = func(values []float64, nsp []uint64) *vector.Vector {
		return makeVector(values, nsp, float64Type)
	}

	MakeCharVector = func(values []string, nsp []uint64) *vector.Vector {
		return makeStringVector(values, nsp, charType)
	}

	MakeVarcharVector = func(values []string, nsp []uint64) *vector.Vector {
		return makeStringVector(values, nsp, varcharType)
	}

	MakeDecimal64Vector = func(values []int64, nsp []uint64, typ types.Type) *vector.Vector {
		vec := vector.New(decimal64Type)
		cols := make([]types.Decimal64, len(values))
		if nsp == nil {
			for i, v := range values {
				cols[i].FromInt64(v)
			}
		} else {
			for _, n := range nsp {
				nulls.Add(vec.Nsp, n)
			}
			for i, v := range values {
				if nulls.Contains(vec.Nsp, uint64(i)) {
					continue
				}
				cols[i].FromInt64(v)
			}
		}
		vec.Col = cols
		return vec
	}

	MakeDecimal128Vector = func(values []int64, nsp []uint64, typ types.Type) *vector.Vector {
		vec := vector.New(decimal128Type)
		cols := make([]types.Decimal128, len(values))
		if nsp == nil {
			for i, v := range values {
				d := types.InitDecimal128(v)
				cols[i] = d
			}
		} else {
			for _, n := range nsp {
				nulls.Add(vec.Nsp, n)
			}
			for i, v := range values {
				if nulls.Contains(vec.Nsp, uint64(i)) {
					continue
				}
				d := types.InitDecimal128(v)
				cols[i] = d
			}
		}
		vec.Col = cols
		return vec
	}

	MakeDateVector = func(values []string, nsp []uint64) *vector.Vector {
		vec := vector.New(dateType)
		ds := make([]types.Date, len(values))
		for _, n := range nsp {
			nulls.Add(vec.Nsp, n)
		}
		for i, s := range values {
			if nulls.Contains(vec.Nsp, uint64(i)) {
				continue
			}
			d, err := types.ParseDate(s)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		vec.Data = encoding.EncodeDateSlice(ds)
		vec.Col = ds
		return vec
	}

	MakeDateTimeVector = func(values []string, nsp []uint64) *vector.Vector {
		vec := vector.New(datetimeType)
		ds := make([]types.Datetime, len(values))
		for _, n := range nsp {
			nulls.Add(vec.Nsp, n)
		}
		for i, s := range values {
			if nulls.Contains(vec.Nsp, uint64(i)) {
				continue
			}
			d, err := types.ParseDatetime(s, 6)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		vec.Data = encoding.EncodeDatetimeSlice(ds)
		vec.Col = ds
		return vec
	}

	MakeTimeStampVector = func(values []string, nsp []uint64) *vector.Vector {
		vec := vector.New(timestampType)
		ds := make([]types.Timestamp, len(values))
		for _, n := range nsp {
			nulls.Add(vec.Nsp, n)
		}
		for i, s := range values {
			if nulls.Contains(vec.Nsp, uint64(i)) {
				continue
			}
			d, err := types.ParseTimestamp(s, 6)
			if err != nil {
				panic(err)
			}
			ds[i] = d
		}
		vec.Data = encoding.EncodeTimestampSlice(ds)
		vec.Col = ds
		return vec
	}
)

// functions to make a scalar vector for test.
var (
	MakeScalarNull = func(length int) *vector.Vector {
		vec := NewProc().AllocScalarNullVector(types.Type{Oid: types.T_any})
		vec.Length = length
		return vec
	}

	MakeScalarBool = func(v bool, length int) *vector.Vector {
		return makeScalar(v, length, boolType)
	}

	MakeScalarInt64 = func(v int64, length int) *vector.Vector {
		return makeScalar(v, length, int64Type)
	}

	MakeScalarInt32 = func(v int32, length int) *vector.Vector {
		return makeScalar(v, length, int32Type)
	}

	MakeScalarInt16 = func(v int16, length int) *vector.Vector {
		return makeScalar(v, length, int16Type)
	}

	MakeScalarInt8 = func(v int8, length int) *vector.Vector {
		return makeScalar(v, length, int8Type)
	}

	MakeScalarUint64 = func(v uint64, length int) *vector.Vector {
		return makeScalar(v, length, uint64Type)
	}

	MakeScalarUint3 = func(v uint32, length int) *vector.Vector {
		return makeScalar(v, length, uint32Type)
	}

	MakeScalarUint16 = func(v uint16, length int) *vector.Vector {
		return makeScalar(v, length, uint16Type)
	}

	MakeScalarUint8 = func(v uint8, length int) *vector.Vector {
		return makeScalar(v, length, uint8Type)
	}

	MakeScalarFloat32 = func(v float32, length int) *vector.Vector {
		return makeScalar(v, length, float32Type)
	}

	MakeScalarFloat64 = func(v float64, length int) *vector.Vector {
		return makeScalar(v, length, float64Type)
	}

	MakeScalarChar = func(value string, length int) *vector.Vector {
		return makeScalarString(value, length, charType)
	}

	MakeScalarVarchar = func(value string, length int) *vector.Vector {
		return makeScalarString(value, length, varcharType)
	}

	MakeScalarDate = func(value string, length int) *vector.Vector {
		vec := NewProc().AllocScalarVector(dateType)
		vec.Length = length
		d, err := types.ParseDate(value)
		if err != nil {
			panic(err)
		}
		vec.Col = []types.Date{d}
		return vec
	}

	MakeScalarDateTime = func(value string, length int) *vector.Vector {
		vec := NewProc().AllocScalarVector(datetimeType)
		vec.Length = length
		d, err := types.ParseDatetime(value, 6)
		if err != nil {
			panic(err)
		}
		vec.Col = []types.Datetime{d}
		return vec
	}

	MakeScalarTimeStamp = func(value string, length int) *vector.Vector {
		vec := NewProc().AllocScalarVector(timestampType)
		vec.Length = length
		d, err := types.ParseTimestamp(value, 6)
		if err != nil {
			panic(err)
		}
		vec.Col = []types.Timestamp{d}
		return vec
	}

	MakeScalarDecimal64 = func(v int64, length int, typ types.Type) *vector.Vector {
		vec := NewProc().AllocScalarVector(decimal64Type)
		vec.Length = length
		vec.Col = []types.Decimal64{types.InitDecimal64(v)}
		return vec
	}

	MakeScalarDecimal128 = func(v uint64, length int, typ types.Type) *vector.Vector {
		vec := NewProc().AllocScalarVector(decimal128Type)
		vec.Length = length
		vec.Col = []types.Decimal128{types.InitDecimal128UsingUint(v)}
		return vec
	}
)

type vecType interface {
	constraints.Integer | constraints.Float | bool
}

func NewProc() *process.Process {
	return process.New(mheap.New(guest.New(1<<20, host.New(1<<20))))
}

func makeVector[T vecType](values []T, nsp []uint64, typ types.Type) *vector.Vector {
	vec := vector.New(typ)
	vec.Data = encoding.EncodeFixedSlice(values, typ.TypeSize())
	vec.Col = values
	for _, n := range nsp {
		nulls.Add(vec.Nsp, n)
	}
	return vec
}

func makeStringVector(values []string, nsp []uint64, typ types.Type) *vector.Vector {
	vec := vector.New(typ)
	bs := &types.Bytes{
		Lengths: make([]uint32, len(values)),
		Offsets: make([]uint32, len(values)),
	}
	next := uint32(0)
	if nsp == nil {
		for i, s := range values {
			l := uint32(len(s))
			bs.Data = append(bs.Data, []byte(s)...)
			bs.Lengths[i] = l
			bs.Offsets[i] = next
			next += l
		}
	} else {
		for _, n := range nsp {
			nulls.Add(vec.Nsp, n)
		}
		for i := range values {
			if nulls.Contains(vec.Nsp, uint64(i)) {
				continue
			}
			s := values[i]
			l := uint32(len(s))
			bs.Data = append(bs.Data, []byte(s)...)
			bs.Lengths[i] = l
			bs.Offsets[i] = next
			next += l
		}
	}
	vec.Col = bs
	return vec
}

func makeScalar[T vecType](value T, length int, typ types.Type) *vector.Vector {
	vec := NewProc().AllocScalarVector(typ)
	vec.Length = length
	vec.Col = []T{value}
	return vec
}

func makeScalarString(value string, length int, typ types.Type) *vector.Vector {
	vec := NewProc().AllocScalarVector(typ)
	vec.Length = length
	l := uint32(len(value))
	vec.Col = &types.Bytes{
		Data:    []byte(value),
		Offsets: []uint32{0},
		Lengths: []uint32{l},
	}
	vec.Data = vec.Col.(*types.Bytes).Data
	return vec
}
