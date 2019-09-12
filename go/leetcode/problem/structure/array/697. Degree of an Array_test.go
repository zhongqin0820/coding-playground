package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/degree-of-an-array/description/

// Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

// Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

func findShortestSubArray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	first := make(map[int]int, n) //第一次出现当前元素的位置
	count := make(map[int]int, n) //记录当前元素出现的个数
	maxCount := 1
	minLen := n
	for i, v := range nums {
		count[v]++
		if count[v] == 1 {
			first[v] = i
		} else {
			l := i - first[v] + 1 //度
			if maxCount < count[v] || (maxCount == count[v] && minLen > l) {
				maxCount = count[v]
				minLen = l
			}
		}
	}

	if len(count) == n {
		return 1
	}
	return minLen
}

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findShortestSubArray(ts.nums))
		})
	}
}
