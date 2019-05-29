package web

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	var c = make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	for v := range c {
		fmt.Print(c)
	}
	close(c)
}
