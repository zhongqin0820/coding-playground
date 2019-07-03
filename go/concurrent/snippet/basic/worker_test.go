package basic

import (
	"log"
	"testing"
	"time"
	"unsafe"
)

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "started job", j)
		time.Sleep(1e9)
		log.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func TestWorker(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w < 4; w++ {
		go workers(w, jobs, results)
	}
	for j := 1; j < 6; j++ {
		jobs <- j
	}
	close(jobs)
	t.Log(unsafe.Sizeof(results), len(results), cap(results))
	for a := 1; a < 6; a++ {
		t.Log(a, <-results)
	}
	t.Log(unsafe.Sizeof(results), len(results), cap(results))
}
