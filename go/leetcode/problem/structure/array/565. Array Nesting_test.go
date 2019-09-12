package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/array-nesting/description/

// A zero-indexed array A of length N contains all integers from 0 to N-1. Find and return the longest length of set S, where S[i] = {A[i], A[A[i]], A[A[A[i]]], ... } subjected to the rule below.

// Suppose the first element in S starts with the selection of element A[i] of index = i, the next element in S should be A[A[i]], and then A[A[A[i]]]… By that analogy, we stop adding right before a duplicate element occurs in S.

func arrayNesting(nums []int) int {
	res := 0
	for i := range nums {
		temp := 1
		// 递归当前nums[i]开始的集合长度
		for nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			temp++
		}

		if res < temp {
			res = temp
		}
	}
	return res
}

func TestArrayNesting(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
		// Explanation:
		// A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.
		// One of the longest S[K]:
		// S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, arrayNesting(ts.nums))
		})
	}
}
