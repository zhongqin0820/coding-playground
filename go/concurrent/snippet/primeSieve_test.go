// A concurrent prime sieve
package snippet

import (
	"log"
	"testing"
)

// ref: https://divan.dev/posts/go_concurrency_visualize/
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func TestPrimeSieve(t *testing.T) {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		log.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
