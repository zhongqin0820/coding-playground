package main

import (
	"log"
)

func generate() chan int {
	ch := make(chan int)
	// defer close(ch)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for prime := range ch {
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func exp1() {
	primes := sieve()
	for {
		log.Println(<-primes)
	}
	return
}

func main() {
	exp1()
}
