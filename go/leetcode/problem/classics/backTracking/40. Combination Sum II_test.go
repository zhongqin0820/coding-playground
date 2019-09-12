package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum-ii/description/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	solution := []int{}
	helper40(candidates, solution, target, &res)
	return res
}

func helper40(candidates, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}
	if len(candidates) == 0 || target < candidates[0] {
		return
	}
	solution = solution[:len(solution):len(solution)]
	helper40(candidates[1:], append(solution, candidates[0]), target-candidates[0], res)
	helper40(helper40Next(candidates), solution, target, res)
}

func helper40Next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, combinationSum2(ts.candidates, ts.target))
		})
	}
}
