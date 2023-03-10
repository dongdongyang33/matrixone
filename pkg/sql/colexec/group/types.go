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

package group

import (
	"github.com/matrixorigin/matrixone/pkg/common/hashmap"
	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/pb/plan"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/agg"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/multi_col/group_concat"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const (
	Build = iota
	BuildWithGroup
	Eval
	EvalWithGroup
	End
)

const (
	H8 = iota
	HStr
)

type evalVector struct {
	needFree bool
	vec      *vector.Vector
}

type container struct {
	typ       int
	state     int
	inserted  []uint8
	zInserted []uint8

	chunkInserted []uint8
	chunkValues   []uint64

	intHashMap *hashmap.IntHashMap
	strHashMap *hashmap.StrHashMap

	aggVecs   []evalVector
	groupVecs []evalVector

	// multiVecs are used for group_concat,
	// cause that group_concat can have many cols like group(a,b,c)
	// in this cases, len(multiVecs[0]) will be 3
	multiVecs [][]evalVector

	vecs []*vector.Vector

	isEmpty bool // Indicates if it is an empty table

	pmIdx int
	pms   []*colexec.PrivMem
}

type Argument struct {
	ctr       *container
	NeedEval  bool // need to projection the aggregate column
	Ibucket   uint64
	Nbucket   uint64
	Exprs     []*plan.Expr // group Expressions
	Types     []types.Type
	Aggs      []agg.Aggregate         // aggregations
	MultiAggs []group_concat.Argument // multiAggs, for now it's group_concat
}

func (ap *Argument) Free(proc *process.Process, pipelineFailed bool) {
	for i := range ap.ctr.pms {
		ap.ctr.pms[i].Clean(proc)
	}
	ap.ctr.cleanHashMap()
}

func (ctr *container) ToInputType(idx int) (t []types.Type) {
	for i := range ctr.multiVecs[idx] {
		t = append(t, ctr.multiVecs[idx][i].vec.Typ)
	}
	return
}

func (ctr *container) ToVecotrs(idx int) (vecs []*vector.Vector) {
	for i := range ctr.multiVecs[idx] {
		vecs = append(vecs, ctr.multiVecs[idx][i].vec)
	}
	return
}

func (ctr *container) cleanAggVectors(mp *mpool.MPool) {
	for i := range ctr.aggVecs {
		if ctr.aggVecs[i].needFree && ctr.aggVecs[i].vec != nil {
			ctr.aggVecs[i].vec.Free(mp)
			ctr.aggVecs[i].vec = nil
		}
	}
}

func (ctr *container) cleanMultiAggVecs(mp *mpool.MPool) {
	for i := range ctr.multiVecs {
		for j := range ctr.multiVecs[i] {
			if ctr.multiVecs[i][j].needFree && ctr.multiVecs[i][j].vec != nil {
				ctr.multiVecs[i][j].vec.Free(mp)
				ctr.multiVecs[i][j].vec = nil
			}
		}
	}
}

func (ctr *container) cleanGroupVectors(mp *mpool.MPool) {
	for i := range ctr.groupVecs {
		if ctr.groupVecs[i].needFree && ctr.groupVecs[i].vec != nil {
			ctr.groupVecs[i].vec.Free(mp)
			ctr.groupVecs[i].vec = nil
		}
	}
}

func (ctr *container) cleanHashMap() {
	if ctr.intHashMap != nil {
		ctr.intHashMap.Free()
		ctr.intHashMap = nil
	}
	if ctr.strHashMap != nil {
		ctr.strHashMap.Free()
		ctr.strHashMap = nil
	}
}
