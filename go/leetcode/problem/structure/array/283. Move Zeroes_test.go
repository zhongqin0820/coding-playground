package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/move-zeroes/description/

// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note:
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.

func moveZeroes(nums []int) {
	j := 0
	for _, v := range nums {
		if v != 0 {
			nums[j] = v
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

func moveZeroesWrapper(nums []int) []int {
	moveZeroes(nums)
	return nums
}

func TestMoveZeroesWrapper(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{1, 0, 1}, []int{1, 1, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, moveZeroesWrapper(ts.nums))
		})
	}
}
