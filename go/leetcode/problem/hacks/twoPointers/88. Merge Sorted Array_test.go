package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, m, n := len(nums1)-1, m-1, n-1; i >= 0; i-- {
		if m >= 0 && n >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else if n >= 0 {
			nums1[i] = nums2[n]
			n--
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			merge(ts.nums1, ts.m, ts.nums2, ts.n)
			ast.Equal(ts.output, ts.nums1)
		})
	}
}
