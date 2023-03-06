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

package minus

import (
	"github.com/matrixorigin/matrixone/pkg/common/hashmap"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

const (
	buildingHashMap = iota
	probingHashMap
	operatorEnd
)

type Argument struct {
	ctr container

	// hash table bucket related information.
	IBucket, NBucket uint64

	// output vector types
	Types []types.Type
}

type container struct {
	colexec.MemforNextOp

	// operator execution stage.
	state int

	// hash table related
	hashTable *hashmap.StrHashMap
}

func (arg *Argument) Free(proc *process.Process, pipelineFailed bool) {
	arg.ctr.CleanMemForNextOp(proc)
	arg.ctr.cleanHashMap()
}

func (ctr *container) cleanHashMap() {
	if ctr.hashTable != nil {
		ctr.hashTable.Free()
		ctr.hashTable = nil
	}
}
