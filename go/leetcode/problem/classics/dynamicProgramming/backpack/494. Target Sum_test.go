package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/target-sum/description/

// You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

// Find out how many ways to assign symbols to make sum of integers equal to target S.

// Note:
// The length of the given array is positive and will not exceed 20.
// The sum of elements in the given array will not exceed 1000.
// Your output answer is guaranteed to be fitted in a 32-bit integer.
func findTargetSumWays(nums []int, S int) int {
	// sum(P)*2 = S+sum(nums)
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// edge check
	if sum < S || (sum+S)&1 == 1 {
		return 0
	}
	// 背包dp模板，切记dp数组的初始化！！
	W := (sum + S) >> 1 //sum(P) = (sum(nums)+S)>>1
	dp := make([]int, W+1)
	dp[0] = 1 //init
	for _, w := range nums {
		for i := W; i >= w; i-- {
			dp[i] = dp[i] + dp[i-w]
		}
	}
	return dp[W]
}

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		S      int
		output int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		// Explanation:
		// -1+1+1+1+1 = 3
		// +1-1+1+1+1 = 3
		// +1+1-1+1+1 = 3
		// +1+1+1-1+1 = 3
		// +1+1+1+1-1 = 3
		// There are 5 ways to assign symbols to make the sum of nums be target 3.
		{[]int{1, 2, 3, 4, 5}, 4, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findTargetSumWays(ts.nums, ts.S))
		})
	}
}
