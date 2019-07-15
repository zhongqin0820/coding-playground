package problem

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	var res [][]int
	helper(&res, []int{}, 1, n, k)
	return res
}

func helper(res *[][]int, coms []int, start int, n int, k int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		*res = append(*res, tmp)
		return
	}
	// backtrack
	for i := start; i <= n; i++ {
		coms = append(coms, i)
		helper(res, coms, i+1, n, k-1)
		coms = coms[:len(coms)-1]
	}
}

func sliceCompare(a, b *[][]int) bool {
	if len((*a)) == 0 || len((*b)) == 0 || len((*a)) != len((*b)) || len((*a)[0]) != len((*b)[0]) {
		return false
	}
	for i, v := range *a {
		for j, val := range (*b)[i] {
			if v[j] != val {
				return false
			}
		}
	}
	return true
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
			if res := combine(ts.n, ts.k); !sliceCompare(&res, &ts.output) {
				t.Errorf("expected %v\n", ts.output)
				t.Logf("got %v\n", res)
			}
		})
	}
}
