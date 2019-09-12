package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-two-integers/description/
func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, (a&b)<<1)
}

func getSumII(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

func TestGetSum(t *testing.T) {
	tests := []struct {
		a, b   int
		output int
	}{
		{1, 2, 3},
		{-2, 3, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getSum(ts.a, ts.b))
			ast.Equal(ts.output, getSumII(ts.a, ts.b))
		})
	}
}
