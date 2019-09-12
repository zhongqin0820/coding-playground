package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/container-with-most-water/

// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

// Note: You may not slant the container and n is at least 2.

func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	i, j, res := 0, len(height)-1, math.MinInt32
	for i < j {
		res = helper11Max(res, helper11Min(height[i], height[j])*(j-i))

		// 朝着area具有变大的可能性方向变化。
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func helper11Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func helper11Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 2, 3, 1}, 3},
		{[]int{1, 1}, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxArea(ts.input))
		})
	}
}
