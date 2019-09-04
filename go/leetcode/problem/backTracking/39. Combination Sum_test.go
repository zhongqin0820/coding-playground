package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(candidates); i++ {
		if target-candidates[i] > 0 {
			for _, res_ := range combinationSum(candidates[i:], target-candidates[i]) {
				temp := make([]int, len(res_)+1)
				copy(temp, res_)
				temp[len(res_)] = candidates[i]
				res = append(res, temp)
			}
		} else if target-candidates[i] == 0 {
			res = append(res, []int{candidates[i]})
		}
	}
	return res
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{3, 2, 2}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {3, 3, 2}, {5, 3}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, combinationSum(ts.candidates, ts.target))
		})
	}
}
