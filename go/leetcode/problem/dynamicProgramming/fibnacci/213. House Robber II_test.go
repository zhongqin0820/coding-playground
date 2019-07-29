package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/house-robber-ii/description/

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

func robII(nums []int) int {
	switch {
	case nums == nil || len(nums) == 0:
		return 0
	case len(nums) == 1:
		return nums[0]
	}
	// 两个区间求最值（0~n-1, 1~n）
	return helper213Max(helper213(nums, 0, len(nums)-1), helper213(nums, 1, len(nums)))
}

func helper213(nums []int, i, j int) int {
	pre, cur := 0, 0
	for k := i; k < j; k++ {
		pre, cur = cur, helper213Max(nums[k]+pre, cur)
	}
	return cur
}

func helper213Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestRobII(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 1, 1, 1}, 3},
		{[]int{2, 3, 2}, 3},
		// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
		//              because they are adjacent houses.
		{[]int{1, 2, 3, 1}, 4},
		// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
		//              Total amount you can rob = 1 + 3 = 4.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, robII(ts.nums))
		})
	}
}
