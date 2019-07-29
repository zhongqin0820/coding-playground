package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum-iv/description/

// Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.

// Note that different sequences are counted as different combinations.
// Follow up:
// What if negative numbers are allowed in the given array?
// How does it change the problem?
// What limitation we need to add to the question to allow negative numbers?

// 涉及顺序的完全背包
func combinationSum4(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	sort.Ints(nums)
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums) && nums[j] <= i; j++ { //物品在内层
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{1, 2, 3}, 4, 7},
		// The possible combination ways are:
		// (1, 1, 1, 1)
		// (1, 1, 2)
		// (1, 2, 1)
		// (1, 3)
		// (2, 1, 1)
		// (2, 2)
		// (3, 1)
		// Therefore the output is 7.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, combinationSum4(ts.nums, ts.target))
		})
	}
}
