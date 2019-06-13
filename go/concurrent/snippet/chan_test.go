package snippet

import (
	"log"
	"testing"
	"time"
)

// ref: https://divan.dev/posts/go_concurrency_visualize/
// fan-in is a function reading from the multiple inputs and multiplexing all into the single channel.
// fan-in = num of inputs
func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		log.Println(x)
	}
}

func TestFanIn(t *testing.T) {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)
	go reader(out)
	for i := range ch {
		out <- i
	}
}

// ref: https://mzh.io/3%E7%A7%8D%E4%BC%98%E9%9B%85%E7%9A%84Go-channel%E7%94%A8%E6%B3%95
func B(g <-chan int, quit <-chan bool) {
	for {
		select {
		case i := <-g:
			log.Println(i + 1)
		case <-quit:
			log.Println("B quit")
			return
		}
	}
}

func TestB(t *testing.T) {
	var g = make(chan int)
	var quit = make(chan bool)
	go B(g, quit)
	// close the chan after writing
	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true // 没办法等待B的退出只能Sleep
	close(g)
	close(quit)
	log.Println("TestB quit")
}
