package snippet

import (
	"log"
	"testing"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			log.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// TestFib is a snippet from https://blog.csdn.net/CodyGuo/article/details/51009768
func TestFib(t *testing.T) {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	t.Logf("\rFib(%d) = %d\n", n, fibN)
}
