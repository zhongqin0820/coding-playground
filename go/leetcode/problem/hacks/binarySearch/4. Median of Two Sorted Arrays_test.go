package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

// There are two sorted arrays nums1 and nums2 of size m and n respectively.

// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

// You may assume nums1 and nums2 cannot be both empty.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return helper4Median(helper4(nums1, nums2))
}

func helper4(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis || (i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs || (i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}
	}
	return res
}

func helper4Median(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}
	return float64(nums[l/2])
}

func TestFindMedianSortedArray(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		output       float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMedianSortedArrays(ts.nums1, ts.nums2))
		})
	}
}
