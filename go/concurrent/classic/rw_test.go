package concurrent

import (
	"strings"
	"sync"
	"testing"
)

func TestNaiveRW(t *testing.T) {
	var m sync.RWMutex
	var rs, ws int
	rsCh := make(chan int)
	wsCh := make(chan int)

	go func() {
		for {
			select {
			case n := <-rsCh:
				rs += n
			case n := <-wsCh:
				ws += n
			}
			t.Logf("%s%s\n", strings.Repeat("R", rs), strings.Repeat("W", ws))
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go NaiveReader(rsCh, &m, &wg)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go NaiveWriter(wsCh, &m, &wg)
	}
	wg.Wait()
}
