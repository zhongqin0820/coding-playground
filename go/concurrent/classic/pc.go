package concurrent

import (
	"log"
	"sync"
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
