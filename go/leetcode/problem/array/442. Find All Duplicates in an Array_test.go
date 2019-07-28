package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?

func findDuplicates(nums []int) []int {
	if nums == nil || len(nums) < 1 {
		return nums
	}
	n := len(nums)
	// 排序令i==nums[i]
	for i := 0; i < n; i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := make([]int, 0, n)
	// 把不在原位的添加进入结果集合
	for i, v := range nums {
		if i != v-1 {
			res = append(res, v)
		}
	}
	return res
}

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{3, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findDuplicates(ts.nums))
		})
	}
}
