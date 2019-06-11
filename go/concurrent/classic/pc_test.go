package classic

import (
	"log"
	"sync"
	"testing"
	"time"
)

type Product struct {
	name string
	id   int
}

func NaiveProducer(c chan Product, wg *sync.WaitGroup, p Product) {
	sleep()
	c <- p
	log.Printf("producer=%s,product=%d\n", p.name, p.id)
	wg.Done()
}

func NaiveConsumer(consumer_ string, c chan Product, wg *sync.WaitGroup) {
	sleep()
	for {
		select {
		case p := <-c:
			log.Printf("consumer=%s,product=%s,%d\n", consumer_, p.name, p.id)
		case <-time.After(time.Second * 3):
			wg.Done()
			return
		}
	}
}

// go test -v -run=TestNaivePC
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
