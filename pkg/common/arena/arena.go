package arena

import (
	"sync/atomic"
	"unsafe"

	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/container/types"
)

// same as golang standard arena, but you don't need to think about concurrency safety
func NewArena(uid uuid.UUID) *Arena {
	return NewArenaWithSize(uid, DefaultArenaSize)
}

// the input unit is Byte
func NewArenaWithSize(uid uuid.UUID, size int) *Arena {
	if size < ChunkSize {
		size = ChunkSize
	}
	chunks := make([]chunk, size/ChunkSize)
	data := make([]byte, ChunkSize*len(chunks))
	for i := range chunks {
		chunks[i] = newChunk(data[i*ChunkSize:])
	}
	return &Arena{
		Uid:    uid,
		data:   data,
		chunks: chunks,
		ptr:    uintptr(unsafe.Pointer(&data[0])),
	}
}

func New[T any](a *Arena) *T {
	var v T

	if sz := round(int(unsafe.Sizeof(v) + WordSize)); sz < PageSize {
		if data := a.alloc(sz); data != nil {
			return (*T)(unsafe.Pointer(&data[0]))
		}
	}
	return new(T)
}

func Free[T any](a *Arena, v *T) {
	a.free(uintptr(unsafe.Pointer(v)))
}

func MakeSlice[T any](a *Arena, len, cap int) []T {
	var v T

	sz := int(unsafe.Sizeof(v))
	if sz := round(sz*cap + WordSize); sz < PageSize {
		if data := a.alloc(sz); data != nil {
			return types.DecodeSlice[T](data)[:len]
		}
	}
	return make([]T, len, cap)
}

func FreeSlice[T any](a *Arena, vs []T) {
	a.free(uintptr(unsafe.Pointer(&vs[0])))
}

func (a *Arena) Free() {
	a.ptr = 0
	a.data = nil
	a.chunks = nil
}

func (a *Arena) alloc(sz int) []byte {
	for i := range a.chunks {
		chunk := &a.chunks[i]
		if data := chunk.alloc(sz); data != nil {
			return data
		}
	}
	return nil
}

func (a *Arena) free(ptr uintptr) {
	if ptr >= a.ptr && ptr < a.ptr+uintptr(len(a.data)) {
		chunk := &a.chunks[(ptr-a.ptr)/ChunkSize]
		chunk.free(ptr)
		return
	}
}

func (ck *chunk) alloc(sz int) []byte {
	for i := range ck.pages {
		pg := &ck.pages[i]
		if slotSize := atomic.LoadInt64(&pg.slotSize); slotSize == -1 {
			for { // ensure that only one caller has exclusive access to this page
				if atomic.CompareAndSwapInt64(&pg.slotSize, -1, 0) {
					pg.init(sz)
					return pg.alloc()
				}
				// failure to seize, abandonment
				if atomic.LoadInt64(&pg.slotSize) != -1 {
					return nil
				}
			}
		} else if slotSize != int64(sz) {
			continue
		}
		if data := pg.alloc(); data != nil {
			return data
		}
	}
	return nil
}

func (ck *chunk) free(ptr uintptr) {
	pg := &ck.pages[(ptr-ck.ptr)/PageSize]
	pg.free(ptr - WordSize)
}

func newChunk(data []byte) chunk {
	pages := make([]page, ChunkSize/PageSize)
	for i := range pages {
		pages[i] = newPage(data[i*PageSize : (i+1)*PageSize])
	}
	return chunk{
		data:  data,
		pages: pages,
		ptr:   uintptr(unsafe.Pointer(&data[0])),
	}
}

func newPage(data []byte) page {
	return page{
		slotSize: -1,
		data:     data,
		ptr:      uintptr(unsafe.Pointer(&data[0])),
	}
}

func (pg *page) init(sz int) {
	atomic.StoreInt64(&pg.head, -1)
	for i := 0; i+sz < PageSize; i += sz {
		atomic.StoreInt64((*int64)(unsafe.Pointer(&pg.data[i])),
			atomic.LoadInt64(&pg.head))
		atomic.StoreInt64(&pg.head, int64(i))
	}
	atomic.StoreInt64(&pg.slotSize, int64(sz))
}

func (pg *page) alloc() []byte {
	for {
		head := atomic.LoadInt64(&pg.head)
		if head == -1 {
			return nil
		}
		if ok := atomic.CompareAndSwapInt64(&pg.head, head,
			atomic.LoadInt64((*int64)(unsafe.Pointer(&pg.data[head])))); ok {
			return pg.data[head+WordSize:]
		}
	}
}

func (pg *page) free(ptr uintptr) {
	ptr -= pg.ptr
	for {
		head := atomic.LoadInt64(&pg.head)
		if ok := atomic.CompareAndSwapInt64(&pg.head, head, int64(ptr)); ok {
			atomic.StoreInt64((*int64)(unsafe.Pointer(&pg.data[ptr])), head)
			return
		}
	}
}

func round(x int) int {
	return ((x + 7) & (-8))
}
