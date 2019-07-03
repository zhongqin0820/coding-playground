package basic

import (
	"sync"
	"testing"
)

func incWithLock(x *int, wg *sync.WaitGroup, l *sync.Mutex) {
	l.Lock()
	*x = *x + 1
	defer l.Unlock()
	defer wg.Done()
}

func incWithChan(x *int, wg *sync.WaitGroup, ch chan struct{}) {
	ch <- struct{}{}
	*x = *x + 1
	<-ch
	defer wg.Done()
}

func BenchmarkInc(b *testing.B) {
	var res1, res2 int
	var lock sync.Mutex
	var wg sync.WaitGroup
	var ch chan struct{} = make(chan struct{})
	//
	b.Run("WithLock", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go incWithLock(&res1, &wg, &lock)
		}
	})
	//
	b.Run("WithChan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go incWithChan(&res2, &wg, ch)
		}
	})
	wg.Wait()
	b.Log(res1)
	b.Log(res2)
}
