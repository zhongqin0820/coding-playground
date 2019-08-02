package snippet

import (
	"log"
	"math"
	"testing"
)

// bit manipulation starts from left to right
func TestBits(t *testing.T) {
	log.Println(4 << ((^uint(0)) >> 63))
	// 4 << 1
	log.Println(1 << 4)

	log.Println(1<<31 - 1)
	log.Println(math.MaxInt32)
	log.Println(1<<63 - 1)
	log.Println(math.MaxInt64)
	log.Println(-1 << 63)
	log.Println(math.MinInt64)
}
