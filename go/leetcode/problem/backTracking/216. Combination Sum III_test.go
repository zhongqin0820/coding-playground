package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/combination-sum-iii/description/
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	temp := make([]int, k+1)
	used := make([]bool, 10)

	var dfs func(int, int)
	dfs = func(idx, remain int) {
		if idx == k {
			if remain < 10 && !used[remain] {
				temp[idx] = remain
				t := make([]int, k)
				copy(t, temp[1:])
				res = append(res, t)
			}
			return
		}

		for i := temp[idx-1] + 1; i < 10; i++ {
			if remain-i < i {
				return
			}
			used[i] = true
			temp[idx] = i
			dfs(idx+1, remain-i)
			used[i] = false
		}
	}

	dfs(1, n)

	return res
}

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		k, n   int
		output [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, combinationSum3(ts.k, ts.n))
		})
	}
}
