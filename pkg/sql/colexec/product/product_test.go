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

package product

import (
	"bytes"
	"context"
	"sync"
	"testing"

	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/hashbuild"
	"github.com/matrixorigin/matrixone/pkg/testutil"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"github.com/stretchr/testify/require"
)

const (
	Rows = 10 // default rows
)

// add unit tests for cases
type productTestCase struct {
	arg    *Argument
	flgs   []bool // flgs[i] == true: nullable
	types  []types.Type
	proc   *process.Process
	cancel context.CancelFunc
	barg   *hashbuild.Argument
}

var (
	tcs []productTestCase
)

func init() {
	tcs = []productTestCase{
		newTestCase([]bool{false}, []types.Type{{Oid: types.T_int8}, {Oid: types.T_int8}}, []colexec.ResultPos{colexec.NewResultPos(0, 0), colexec.NewResultPos(1, 0)}),
		newTestCase([]bool{true}, []types.Type{{Oid: types.T_int8}, {Oid: types.T_int8}}, []colexec.ResultPos{colexec.NewResultPos(0, 0), colexec.NewResultPos(1, 0)}),
	}
}

func TestString(t *testing.T) {
	buf := new(bytes.Buffer)
	for _, tc := range tcs {
		String(tc.arg, buf)
	}
}

func TestProduct(t *testing.T) {
	for _, tc := range tcs {
		bat := hashBuild(t, tc)
		var wg sync.WaitGroup

		err := Prepare(tc.proc, tc.arg)
		require.NoError(t, err)
		wg.Add(1)
		go func() {
			for {
				if ok, err := Call(0, tc.proc, tc.arg, false, false); ok || err != nil {
					wg.Done()
					break
				}
			}
		}()
		tc.proc.Reg.MergeReceivers[1].Ch <- bat
		bat0 := newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
		bat0.AddCnt(1)
		tc.proc.Reg.MergeReceivers[0].Ch <- bat0
		tc.proc.Reg.MergeReceivers[0].Ch <- &batch.Batch{}
		bat1 := newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
		bat1.AddCnt(1)
		tc.proc.Reg.MergeReceivers[0].Ch <- bat1
		bat2 := newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
		bat2.AddCnt(1)
		tc.proc.Reg.MergeReceivers[0].Ch <- bat2
		bat3 := newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
		bat3.AddCnt(1)
		tc.proc.Reg.MergeReceivers[0].Ch <- bat3
		tc.proc.Reg.MergeReceivers[0].Ch <- nil
		wg.Wait()
		bat0.Clean(tc.proc.Mp())
		bat1.Clean(tc.proc.Mp())
		bat2.Clean(tc.proc.Mp())
		bat3.Clean(tc.proc.Mp())
		tc.arg.Free(tc.proc, false)
		tc.barg.Free(tc.proc, false)
		require.Equal(t, int64(0), tc.proc.Mp().CurrNB())
		/*
			tc.proc.Reg.MergeReceivers[0].Ch <- newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
			tc.proc.Reg.MergeReceivers[0].Ch <- &batch.Batch{}
			tc.proc.Reg.MergeReceivers[0].Ch <- newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
			tc.proc.Reg.MergeReceivers[0].Ch <- newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
			tc.proc.Reg.MergeReceivers[0].Ch <- newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
			tc.proc.Reg.MergeReceivers[0].Ch <- nil
			tc.proc.Reg.MergeReceivers[1].Ch <- bat
			for {
				if ok, err := Call(0, tc.proc, tc.arg, false, false); ok || err != nil {
					break
				}
				tc.proc.Reg.InputBatch.Clean(tc.proc.Mp())
			}
			require.Equal(t, int64(0), tc.proc.Mp().CurrNB())
		*/
	}
}

func newTestCase(flgs []bool, ts []types.Type, rp []colexec.ResultPos) productTestCase {
	ctx, cancel := context.WithCancel(context.Background())
	proc := process.NewFromProc(testutil.NewProcessWithMPool(mpool.MustNewZero()),
		ctx, -1)
	return productTestCase{
		types:  ts,
		flgs:   flgs,
		proc:   proc,
		cancel: cancel,
		arg: &Argument{
			Types:  ts,
			Result: rp,
		},
		barg: &hashbuild.Argument{
			Types: ts,
		},
	}
}

func hashBuild(t *testing.T, tc productTestCase) *batch.Batch {
	var wg sync.WaitGroup

	err := hashbuild.Prepare(tc.proc, tc.barg)
	require.NoError(t, err)
	wg.Add(1)
	go func() {
		for {
			if ok, err := hashbuild.Call(0, tc.proc, tc.barg, false, false); ok || err != nil {
				wg.Done()
				break
			}
		}
	}()
	bat := newBatch(t, tc.flgs, tc.types, tc.proc, Rows)
	bat.AddCnt(1)
	tc.proc.Reg.MergeReceivers[0].Ch <- bat
	tc.proc.Reg.MergeReceivers[0].Ch <- nil
	bat.Clean(tc.proc.Mp())
	tc.proc.InputBatch().AddCnt(1)
	return tc.proc.InputBatch()
}

// create a new block based on the type information, flgs[i] == ture: has null
func newBatch(t *testing.T, flgs []bool, ts []types.Type, proc *process.Process, rows int64) *batch.Batch {
	return testutil.NewBatch(ts, false, int(rows), proc.Mp())
}
