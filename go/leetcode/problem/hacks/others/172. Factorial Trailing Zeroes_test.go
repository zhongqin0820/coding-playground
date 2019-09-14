package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/factorial-trailing-zeroes/description/
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		n /= 5
		res += n
	}
	return res
}

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{3, 0},
		{5, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, trailingZeroes(ts.n))
		})
	}
}
