package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// nums[mid] == target
			// 移动左指针
			if nums[left] < target {
				left++
			}
			// 移动右指针
			if nums[right] > target {
				right--
			}
			// 返回结果
			if nums[left] == nums[right] {
				return []int{left, right}
			}
		}
	}
	return []int{-1, -1}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{1, 2}, 1, []int{0, 0}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, searchRange(ts.nums, ts.target))
		})
	}
}
