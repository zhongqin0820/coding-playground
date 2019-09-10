package snippet

import (
	"log"
	"math"
	"math/bits"
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

func TestMathBits(t *testing.T) {
	var x, y uint = 1, 2 << 31
	log.Println(bits.Mul(x, y))
	var a uint = 2<<31 - 2<<28
	log.Printf("%032b %032b = %032b, %d", 2<<31, 2<<28, a, bits.OnesCount(a))
	var b uint = 2 << 31
	log.Printf("%032b %032b", b, bits.Reverse(b))
	log.Println(bits.Len(uint(8)))
}
