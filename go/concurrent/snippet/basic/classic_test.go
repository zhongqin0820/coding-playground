package basic

import (
	"log"
	"sync"
	"testing"
)

func Do() <-chan int {
	var a chan int = make(chan int)
	i := 0
	go func() {
		for ; i < 5; i++ {
			a <- i
		}
		close(a)
	}()
	return a
}

func TestClassic(t *testing.T) {
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for v := range Do() {
			log.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
	t.Log()
}
