package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	sum, res := -1<<31, -1<<31
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, n := range nums {
		sum = max(sum+n, n)
		res = max(res, sum)
	}
	return res
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxSubArray(ts.nums))
		})
	}
}
