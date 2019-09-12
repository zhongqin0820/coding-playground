package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/single-element-in-a-sorted-array/description/
// O(n)
func singleNonDuplicate(nums []int) int {
	res, n := nums[0], len(nums)
	for i := 1; i < n; i++ {
		res ^= nums[i]
	}
	return res
}

// O(log(n))
func naiveSingleNonDuplicate(nums []int) int {
	n := len(nums)
	l, h := 0, n-1
	for l < h {
		m := l + (h-l)/2
		// 确保m在偶数位上
		if m%2 == 1 {
			m--
		}
		// 判断条件
		if nums[m] == nums[m+1] {
			l = m + 2
		} else {
			h = m
		}
	}
	return nums[l]
}

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, singleNonDuplicate(ts.nums))
			ast.Equal(ts.output, naiveSingleNonDuplicate(ts.nums))
		})
	}
}
