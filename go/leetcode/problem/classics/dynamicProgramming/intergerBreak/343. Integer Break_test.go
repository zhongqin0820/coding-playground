package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/integer-break/description/

// Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

// Note: You may assume that n is not less than 2 and not larger than 58.
// ref: https://leetcode.com/problems/integer-break/discuss/80689/A-simple-explanation-of-the-math-part-and-a-O(n)-solution
func integerBreak(n int) int {
	switch n {
	case 2, 3:
		return n - 1
	}
	switch n % 3 {
	case 0:
		return helper343(n / 3)
	case 1:
		return helper343(n/3-1) * 4
	default:
		return helper343(n/3) * 2
	}
}

func helper343(n int) int {
	if n == 0 {
		return 1
	}
	res := helper343(n >> 1)
	if n&1 == 0 {
		return res * res
	}
	return res * res * 3
}

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{2, 1},
		// Explanation: 2 = 1 + 1, 1 × 1 = 1.
		{10, 36},
		// Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, integerBreak(ts.n))
		})
	}
}
