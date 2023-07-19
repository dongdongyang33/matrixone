package arena

import (
	"sync"
	"testing"
)

// test race
func TestArenaForRace(t *testing.T) {
	var wg sync.WaitGroup

	cl := NewArena()
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
