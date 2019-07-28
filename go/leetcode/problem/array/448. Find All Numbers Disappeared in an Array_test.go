package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements of [1, n] inclusive that do not appear in this array.
// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

func findDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, len(nums))
	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findDisappearedNumbers(ts.nums))
		})
	}
}
