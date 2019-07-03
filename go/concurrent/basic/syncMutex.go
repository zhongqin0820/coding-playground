package main

import "sync"

// sync
// go test -race
// sync.Mutex
// sync.RWMutex
// sync with mutex
type Counter struct {
	sync.Mutex //embedding
	value      int
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			c.Lock()
			defer c.Unlock()
			defer wg.Done()
			c.value++
		}()
	}
	wg.Wait()
	println(c.value)
}
