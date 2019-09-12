package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
func findMin(nums []int) int {
	n := len(nums)
	l, h := 0, n-1
	for l < h {
		m := l + (h-l)/2
		// 右边都是排好序的
		if nums[m] < nums[h] {
			h = m
		} else {
			// 左边都是排好序的
			l = m + 1
		}
	}
	return nums[l]
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMin(ts.nums))
		})
	}
}
