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

package disttae

import (
	"fmt"
	"testing"

	"github.com/matrixorigin/matrixone/pkg/pb/api"
	"github.com/matrixorigin/matrixone/pkg/pb/logtail"
	"github.com/matrixorigin/matrixone/pkg/pb/timestamp"
	"github.com/stretchr/testify/require"
)

// the sort func should ensure that:
// the smaller the table id of item, the smaller the indexing.
// the smaller the timestamp of item, the smaller the indexing.
// if timestamps are equal, the smaller table id one will be in front (with a smaller index).
func TestLogList_Sort(t *testing.T) {
	var lists logList = []logtail.TableLogtail{
		{Ts: &timestamp.Timestamp{PhysicalTime: 1}, Table: &api.TableID{TbId: 2}},
		{Ts: &timestamp.Timestamp{PhysicalTime: 1}, Table: &api.TableID{TbId: 1}},
		{Ts: &timestamp.Timestamp{PhysicalTime: 1}, Table: &api.TableID{TbId: 3}},
		{Ts: &timestamp.Timestamp{PhysicalTime: 2}, Table: &api.TableID{TbId: 1}},
		{Ts: &timestamp.Timestamp{PhysicalTime: 3}, Table: &api.TableID{TbId: 1}},
	}
	lists.Sort()
	expectedT := []int64{1, 1, 1, 2, 3}
	expectedI := []uint64{1, 2, 3, 1, 1}
	for i := range lists {
		require.Equal(t, expectedT[i], lists[i].Ts.PhysicalTime)
		require.Equal(t, expectedI[i], lists[i].Table.TbId)
	}
}

/*
// should ensure syncLogTailTimestamp can get time or set time without any trace.
// and will not cause any deadlock.
func TestSyncLogTailTimestamp(t *testing.T) {
	// Not now.
	t.Skip()
}
*/

// should ensure that subscribe and unsubscribe methods are effective.
func TestSubscribedTable(t *testing.T) {
	var subscribeRecord subscribedTable
	subscribeRecord.initTableSubscribeRecord()
	require.Equal(t, 0, len(subscribeRecord.m))

	tbls := []struct {
		db uint64
		tb uint64
	}{
		{db: 0, tb: 1},
		{db: 0, tb: 2},
		{db: 0, tb: 3},
		{db: 0, tb: 4},
		{db: 0, tb: 1},
	}
	for _, tbl := range tbls {
		subscribeRecord.setTableSubscribe(tbl.db, tbl.tb)
	}
	require.Equal(t, 4, len(subscribeRecord.m))
	require.Equal(t, true, subscribeRecord.getTableSubscribe(tbls[0].db, tbls[0].tb))
	for _, tbl := range tbls {
		subscribeRecord.setTableUnsubscribe(tbl.db, tbl.tb)
	}
	require.Equal(t, 0, len(subscribeRecord.m))
}

var _ = debugToPrintLogList

func debugToPrintLogList(ls []logtail.TableLogtail) string {
	if len(ls) == 0 {
		return ""
	}
	str := "log list are:\n"
	for i, l := range ls {
		did, tid := l.Table.DbId, l.Table.TbId
		str += fmt.Sprintf("\t log: %d, dn: %d, tbl: %d\n", i, did, tid)
		if len(l.Commands) > 0 {
			str += "\tcommands are :\n"
		}
		for j, command := range l.Commands {
			typ := "insert"
			if command.EntryType == 1 {
				typ = "delete"
			} else if command.EntryType == 2 {
				typ = "update"
			}
			str += fmt.Sprintf("\t\t %d: [dnName: %s, tableName: %s, typ: %s]\n",
				j, command.DatabaseName, command.TableName, typ)
		}
	}
	return str
}
