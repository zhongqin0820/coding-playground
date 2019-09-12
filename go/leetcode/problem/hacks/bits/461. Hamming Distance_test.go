package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	// 对异或后的结果统计1的个数
	x ^= y
	res := 0
	for x > 0 {
		res += (x & 1)
		x >>= 1
	}
	return res
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x, y   int
		output int
	}{
		{1, 4, 2},
		{3, 1, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, hammingDistance(ts.x, ts.y))
		})
	}
}
