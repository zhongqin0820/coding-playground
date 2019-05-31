package concurrent

import (
	"sync"
	"testing"
)

func TestNaivePC(t *testing.T) {
	var wg = sync.WaitGroup{}
	var c = make(chan Product, 10)
	var producers = []string{"a", "b"}
	var consumers = []string{"A", "B"}
	var ids = []int{1, 2}
	// producer
	for _, producer_ := range producers {
		for _, id_ := range ids {
			wg.Add(1)
			p_ := Product{name: producer_, id: id_}
			go func(p Product) {
				NaiveProducer(c, &wg, p)
			}(p_)
		}
	}

	// consumer
	for _, consumer_ := range consumers {
		wg.Add(1)
		go NaiveConsumer(consumer_, c, &wg)

	}

	wg.Wait()
	close(c)
}
