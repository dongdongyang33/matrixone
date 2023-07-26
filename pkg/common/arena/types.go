package arena

import "github.com/google/uuid"

const (
	// 4k may be a good size for page
	PageSize = 4 << 10
	// 1MB may be a good size for chunk
	ChunkSize = 1 << 20
	WordSize  = 8

	DefaultArenaSize = 100 << 20
)

type page struct {
	slotSize int64
	head     int64
	data     []byte
	// start pointer
	ptr uintptr
}

type chunk struct {
	pages []page
	data  []byte
	// start pointer
	ptr uintptr
}

type Arena struct {
	Cnt    int64
	Uid    uuid.UUID
	data   []byte
	chunks []chunk
	// start pointer
	ptr uintptr
}
