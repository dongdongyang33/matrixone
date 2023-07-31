package arena

import (
	"fmt"
	"sync"
	"testing"

	"github.com/google/uuid"
)

type MyByte [7]byte

// test race
func TestArenaForRace(t *testing.T) {
	var wg sync.WaitGroup

	cl := NewArena(uuid.New())
	run := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			v := New[int64](cl)
			Free(cl, v)
		}
	}
	for i := 0; i < 800; i++ {
		wg.Add(1)
		go run()
	}
	wg.Wait()
	cl.Free()
}

func TestArenaForRace1(t *testing.T) {

	cl := NewArena(uuid.New())
	dd := MakeSlice[MyByte](cl, 1, 1)

	fmt.Printf("dd len = %d, cap =%d", len(dd), cap(dd))
}
