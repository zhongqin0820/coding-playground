package problem

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/power-of-two/description/
func isPowerOfTwo(n int) bool {
	return n > 0 && bits.OnesCount64(uint64(n)) == 1
}

func isPowerOfTwoI(n int) bool {
	// 利用与运算性质：1000 & 0111 = 0
	return n > 0 && (n&(n-1)) == 0
}

func isPowerOfTwoII(n int) bool {
	if n <= 0 {
		return false
	}
	res := 0
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res++
			if res > 1 {
				return false
			}
		}
	}
	return true
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n      int
		output bool
	}{
		{1, true},
		{16, true},
		{218, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPowerOfTwo(ts.n))
			ast.Equal(ts.output, isPowerOfTwoI(ts.n))
			ast.Equal(ts.output, isPowerOfTwoII(ts.n))
		})
	}
}
