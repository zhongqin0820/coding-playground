package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subsets-ii/

// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

func subsetsWithDup(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{[]int{}}
	var cur []int
	sort.Ints(nums)
	helper90(nums, cur, &res)
	return res
}

func helper90(nums []int, cur []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		// skip the same number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// backtracking
		cur = append(cur, nums[i])

		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)

		helper90(nums[i+1:], cur, res)
		cur = cur[:len(cur)-1]
	}
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{nil, [][]int{}},
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, subsetsWithDup(ts.input))
		})
	}

}
