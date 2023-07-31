package arena

import (
	"sync"

	"github.com/google/uuid"
)

const (
	// 4k may be a good size for page
	PageSize = 4 << 10
	// 1MB may be a good size for chunk
	ChunkSize = 1 << 20
	WordSize  = 8

	DefaultArenaSize = 50 << 20
)

type page struct {
	slotSize int64
	data     []byte
	ptr      uintptr // start pointer

	head int64
}

type chunk struct {
	pages []page
	data  []byte
	// start pointer
	ptr uintptr
}

type Arena struct {
	sync.Mutex
	Cnt int64
	Uid uuid.UUID

	data   []byte
	chunks []chunk
	// start pointer
	ptr uintptr

	mp map[uintptr]bool
}
