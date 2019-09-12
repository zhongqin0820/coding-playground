package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subsets/description/

func subsets(nums []int) [][]int {
	res := make([][]int, 1, 1024)
	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums   []int
		output [][]int
	}{
		{[]int{1, 2, 3}, [][]int{nil, {1}, {2}, {2, 1}, {3}, {3, 1}, {3, 2}, {3, 2, 1}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, subsets(ts.nums))
		})
	}
}
