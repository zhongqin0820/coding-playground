package problem

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/combinations/

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

func combine(n int, k int) [][]int {
	var res [][]int
	helper77(&res, []int{}, 1, n, k)
	return res
}

func helper77(res *[][]int, coms []int, start int, n int, k int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		*res = append(*res, tmp)
		return
	}
	// backtrack
	for i := start; i <= n; i++ {
		coms = append(coms, i)
		helper77(res, coms, i+1, n, k-1)
		coms = coms[:len(coms)-1]
	}
}

func TestCombine(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		output [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			if res := combine(ts.n, ts.k); !slice2DCompare(&res, &ts.output) {
				t.Errorf("expected %v\n", ts.output)
				t.Errorf("got %v\n", res)
			}
		})
	}
}
