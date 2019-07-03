package main

import "sync"

func main() {
	c := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		defer wg.Done()
		for v := range c {
			println(v)
		}
	}()
	wg.Wait()
}
