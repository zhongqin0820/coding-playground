package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/partition-equal-subset-sum/description/

// Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

// Note:
// Each of the array element will not exceed 100.
// The array size will not exceed 200.

func canPartition(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	W := sum / 2
	dp := make([]bool, W+1)
	dp[0] = true
	for _, num := range nums {
		for i := W; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[W]
}

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums   []int
		output bool
	}{
		{[]int{}, true},
		{[]int{1, 5, 11, 5}, true},
		// Explanation: The array can be partitioned as [1, 5, 5] and [11].
		{[]int{1, 2, 3, 5}, false},
		// Explanation: The array cannot be partitioned into equal sum subsets.
		{[]int{2, 2, 3, 5}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canPartition(ts.nums))
		})
	}
}
