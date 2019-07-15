package problem

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		// 判断是否找到
		if nums[m] == target {
			return true
		}
		//
		if nums[i] == nums[m] {
			i++
		} else if nums[i] < nums[m] {
			// 说明左侧递增
			// target 属于 [nums[i], nums[m])，向左侧缩小范围
			if nums[i] <= target && target < nums[m] {
				j = m - 1
			} else {
				i = m + 1
			}
		} else {
			// 说明右侧递增
			// target 属于 (nums[m], nums[j]]，向右侧缩小范围
			if nums[m] < target && target <= nums[j] {
				i = m + 1
			} else {
				j = m - 1
			}
		}
	}
	return false
}

func searchAdv(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	// 找到pivot
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}
	// 针对pivot进行调整范围
	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return true
		}
	}

	return false
}

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output bool
	}{
		{nil, 0, false},
		{[]int{}, 0, false},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{3, 1, 2, 3, 3, 3, 3}, 1, true},
		{[]int{1}, 1, true},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			if res := search(ts.nums, ts.target); res != ts.output {
				t.Errorf("expected %t got %t\n", ts.output, res)
			}
		})
		t.Run(fmt.Sprintf("Example %d Adv", i+1), func(t *testing.T) {
			if res := searchAdv(ts.nums, ts.target); res != ts.output {
				t.Errorf("expected %t got %t\n", ts.output, res)
			}
		})
	}
}
