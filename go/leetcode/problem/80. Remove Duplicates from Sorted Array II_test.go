package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

// Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i, j, n := 1, 2, len(nums)
	for ; j < n; j++ {
		// [i-1]!=[i] || [i]!=[j]
		if nums[i] != nums[j] || nums[i] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func TestRemoveDuplications(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		// normal test
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		// edge test
		{nil, 0},
		{[]int{}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, removeDuplicates(ts.input))
		})
	}
}
