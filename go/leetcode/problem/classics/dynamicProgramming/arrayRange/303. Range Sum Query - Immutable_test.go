package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/range-sum-query-immutable/description/

// Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

// Note:
// You may assume that **the array does not change.**
// There are many calls to sumRange function.

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return NumArray{dp: dp}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j+1] - this.dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func TestNumArray(t *testing.T) {
	// Example:
	ast := assert.New(t)
	nums := []int{-2, 0, 3, -5, 2, -1}
	// Given nums = [-2, 0, 3, -5, 2, -1]
	obj := Constructor(nums)
	ast.Equal(-2, obj.SumRange(0, 0))
	ast.Equal(1, obj.SumRange(0, 2))
	// sumRange(0, 2) -> 1
	ast.Equal(-1, obj.SumRange(2, 5))
	// sumRange(2, 5) -> -1
	ast.Equal(-3, obj.SumRange(0, 5))
	// sumRange(0, 5) -> -3
}
