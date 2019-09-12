package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/house-robber/description/

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return helper198(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], helper198(nums[0], nums[1])
	res := helper198(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = helper198(dp[i-1], nums[i]+dp[i-2])
		res = helper198(res, dp[i])
	}
	return res
}

func robAdv(nums []int) int {
	if nums == nil {
		return 0
	}
	pre, res, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		pre, res = res, helper198(nums[i]+pre, res)
	}
	return res
}

func helper198(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestRob(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3, 1}, 4},
		// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
		//              Total amount you can rob = 1 + 3 = 4.
		{[]int{2, 7, 9, 3, 1}, 12},
		// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
		//              Total amount you can rob = 2 + 9 + 1 = 12.

	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, rob(ts.nums))
			ast.Equal(ts.output, robAdv(ts.nums))
		})
	}
}
