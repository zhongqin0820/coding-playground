package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var baton = make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {

	var newRunner int
	runner, ok := <-baton
	if !ok {
		fmt.Println("Chan Closed!")
		return
	}
	fmt.Printf("Runner %d run\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d come to line\n", newRunner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d finished, Done\n", runner)
		close(baton)
		wg.Done()
		return
	}
	fmt.Printf("%d exchanged to %d\n", runner, newRunner)
	baton <- newRunner
}
