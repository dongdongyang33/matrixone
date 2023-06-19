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

package colexec

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/common/moerr"
	"github.com/matrixorigin/matrixone/pkg/logservice"
	"github.com/matrixorigin/matrixone/pkg/objectio"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

var Srv *Server

const (
	TxnWorkSpaceIdType = 1
	CnBlockIdType      = 2
)

func NewServer(client logservice.CNHAKeeperClient) *Server {
	if Srv != nil {
		return Srv
	}
	Srv = &Server{
		mp:       make(map[uint64]*process.WaitRegister),
		hakeeper: client,
		uuidCsChanMap: UuidProcMap{
			cntMp: make(map[uuid.UUID]int),
			mp:    make(map[uuid.UUID][]*process.Process),
		},
		cnSegmentMap: CnSegmentMap{mp: make(map[objectio.Segmentid]int32)},
	}
	return Srv
}

func (srv *Server) GetConnector(id uint64) *process.WaitRegister {
	srv.Lock()
	defer srv.Unlock()
	defer func() { delete(srv.mp, id) }()
	return srv.mp[id]
}

func (srv *Server) RegistConnector(reg *process.WaitRegister) uint64 {
	srv.Lock()
	defer srv.Unlock()
	srv.mp[srv.id] = reg
	defer func() { srv.id++ }()
	return srv.id
}

func (srv *Server) PutParallelNumForUuid(u uuid.UUID, num int) error {
	srv.uuidCsChanMap.Lock()
	defer srv.uuidCsChanMap.Unlock()
	if n, ok := srv.uuidCsChanMap.cntMp[u]; ok {
		if n != num {
			fmt.Printf("[putuuidnum] error. uuid = %s, old = %d, new = %d\n", u, n, num)
			return moerr.NewInternalErrorNoCtx("uuid %s parallel count had been set to %d\n", u, n)
		}
		return nil
	}
	fmt.Printf("[putuuidnum] uuid %s -> num %d\n", u, num)
	srv.uuidCsChanMap.cntMp[u] = num
	return nil
}

func (srv *Server) GetParallelNumForUuid(u uuid.UUID) (int, bool) {
	srv.uuidCsChanMap.Lock()
	defer srv.uuidCsChanMap.Unlock()

	ret, ok := srv.uuidCsChanMap.cntMp[u]
	return ret, ok
}

func (srv *Server) GetProcByUuid(u uuid.UUID, idx int) (*process.Process, bool) {
	srv.uuidCsChanMap.Lock()
	defer srv.uuidCsChanMap.Unlock()

	ps, ok := srv.uuidCsChanMap.mp[u]
	if !ok {
		return nil, false
	}

	if idx >= len(ps) {
		return nil, false
	}

	return ps[idx], true
}

func (srv *Server) PutProcIntoUuidMap(u uuid.UUID, p *process.Process) error {
	srv.uuidCsChanMap.Lock()
	defer srv.uuidCsChanMap.Unlock()
	ps, ok := srv.uuidCsChanMap.mp[u]
	if ok {
		ps = append(ps, p)
		srv.uuidCsChanMap.mp[u] = ps
		fmt.Printf("[putuuid] put uuid %s with proc %p success, current len = %d\n", u, p, len(ps))
	} else {
		srv.uuidCsChanMap.mp[u] = []*process.Process{p}
		fmt.Printf("[putuuid] (first) put uuid %s with proc %p success, current len = %d\n", u, p, len(srv.uuidCsChanMap.mp[u]))
	}
	return nil
}

func (srv *Server) CleanUuidFromMap(u uuid.UUID) error {
	srv.uuidCsChanMap.Lock()
	defer srv.uuidCsChanMap.Unlock()
	fmt.Printf("[cleanuuid] clean uuid %s\n", u)
	delete(srv.uuidCsChanMap.cntMp, u)
	delete(srv.uuidCsChanMap.mp, u)
	return nil
}

func (srv *Server) PutCnSegment(segmentName *objectio.Segmentid, segmentType int32) {
	srv.cnSegmentMap.Lock()
	defer srv.cnSegmentMap.Unlock()
	srv.cnSegmentMap.mp[*segmentName] = segmentType
}

func (srv *Server) DeleteTxnSegmentIds(segmentNames []objectio.Segmentid) {
	srv.cnSegmentMap.Lock()
	defer srv.cnSegmentMap.Unlock()
	for _, segmentName := range segmentNames {
		delete(srv.cnSegmentMap.mp, segmentName)
	}
}

func (srv *Server) GetCnSegmentMap() map[string]int32 {
	srv.cnSegmentMap.Lock()
	defer srv.cnSegmentMap.Unlock()
	new_mp := make(map[string]int32)
	for k, v := range srv.cnSegmentMap.mp {
		new_mp[string(k[:])] = v
	}
	return new_mp
}

func (srv *Server) GetCnSegmentType(segmentName *objectio.Segmentid) int32 {
	srv.cnSegmentMap.Lock()
	defer srv.cnSegmentMap.Unlock()
	return srv.cnSegmentMap.mp[*segmentName]
}

// SegmentId is part of Id for cn2s3 directly, for more info, refer to docs about it
func (srv *Server) GenerateSegment() objectio.ObjectName {
	srv.Lock()
	defer srv.Unlock()
	return objectio.BuildObjectName(objectio.NewSegmentid(), 0)
	// for future fileOffset
	// if srv.InitSegmentId {
	// 	srv.incrementSegmentId()
	// } else {
	// 	srv.getNewSegmentId()
	// 	srv.currentFileOffset = 0
	// 	srv.InitSegmentId = true
	// }
	// return objectio.BuildObjectName(srv.CNSegmentId, srv.currentFileOffset)
}

// func (srv *Server) incrementSegmentId() {
// 	if srv.currentFileOffset < math.MaxUint16 {
// 		srv.currentFileOffset++
// 	} else {
// 		srv.getNewSegmentId()
// 		srv.currentFileOffset = 0
// 	}
// }

// // for now, rowId is common between CN and DN.
// func (srv *Server) getNewSegmentId() {
// 	srv.CNSegmentId = common.NewSegmentid()
// }
