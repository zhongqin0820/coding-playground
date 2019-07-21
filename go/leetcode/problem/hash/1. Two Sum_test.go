package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum/

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 1 {
		return []int{}
	}
	res := []int{}
	m := map[int]int{}
	for j, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = j
	}
	return res
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{nil, 0, []int{}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 0, []int{}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, -1, []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, twoSum(ts.nums, ts.target))
		})
	}
}
