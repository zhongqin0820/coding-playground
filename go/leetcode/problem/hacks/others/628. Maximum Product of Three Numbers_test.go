package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-of-three-numbers/description/
func maximumProduct(nums []int) int {
	maxF := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 可能的组合情况：一个正数，（两个正数，连个负数）
	max, max1, max2, min1, min2 := -1001, -1001, -1001, 1001, 1001
	for _, n := range nums {
		switch {
		case n > max:
			max2, max1, max = max1, max, n
		case n > max1:
			max2, max1 = max1, n
		case n > max2:
			max2 = n
		}

		switch {
		case n < min1:
			min2, min1 = min1, n
		case n < min2:
			min2 = n
		}
	}

	return maxF(max1*max2, min1*min2) * max
}

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{-4, -3, -2, -1, 1, 1, 1, 1, 1, 60}, 720},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maximumProduct(ts.nums))
		})
	}
}
