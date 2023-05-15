// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"github.com/matrixorigin/matrixone/pkg/catalog"
	"github.com/matrixorigin/matrixone/pkg/pb/plan"
	"github.com/matrixorigin/matrixone/pkg/sql/parsers/tree"
)

var (
	MetadataScanColDefs = []*plan.ColDef{
		{
			Name: catalog.MetadataScanInfoNames[catalog.COL_NAME],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.COL_NAME].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.BLOCK_ID],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.BLOCK_ID].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.ENTRY_STATE],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.ENTRY_STATE].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.SORTED],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.SORTED].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.META_LOC],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.META_LOC].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.DELTA_LOC],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.DELTA_LOC].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.COMMIT_TS],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.COMMIT_TS].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.META_SEG],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.META_SEG].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.ROW_CNT],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.ROW_CNT].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.NULL_CNT],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.NULL_CNT].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.COMPRESS_SIZE],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.COMPRESS_SIZE].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.ORIGIN_SIZE],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.ORIGIN_SIZE].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.MIN],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.MIN].Oid),
				NotNullable: false,
			},
		},
		{
			Name: catalog.MetadataScanInfoNames[catalog.MAX],
			Typ: &plan.Type{
				Id:          int32(catalog.MetaColTypes[catalog.MAX].Oid),
				NotNullable: false,
			},
		},
	}
)

func (builder *QueryBuilder) buildMetadataScan(tbl *tree.TableFunction, ctx *BindContext, exprs []*plan.Expr, childId int32) int32 {
	node := &plan.Node{
		NodeType: plan.Node_FUNCTION_SCAN,
		Stats:    &plan.Stats{},
		TableDef: &plan.TableDef{
			TableType: "func_table",
			TblFunc: &plan.TableFunction{
				Name: "metadata_scan",
			},
			Cols: MetadataScanColDefs,
		},
		BindingTags:     []int32{builder.genNewTag()},
		Children:        []int32{childId},
		TblFuncExprList: exprs,
	}
	return builder.appendNode(node, ctx)
}
