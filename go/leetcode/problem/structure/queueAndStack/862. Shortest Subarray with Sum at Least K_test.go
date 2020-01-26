package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/
// ref: https://www.cnblogs.com/grandyang/p/11300071.html
func shortestSubarray(A []int, K int) int {
	size, res := len(A), 1<<31
	dq := make([]int, 0, size)
	// 连续子序列和数组，sums[i,j]==sums[j]-sums[i]
	sums := make([]int, size+1)
	for i := 1; i <= size; i++ {
		sums[i] = sums[i-1] + A[i-1]
	}
	for i := 0; i <= size; i++ {
		// 符合条件的i
		for len(dq) != 0 && sums[i]-sums[dq[0]] >= K {
			res = helper862Min(res, i-dq[0])
			dq = dq[1:]
		}
		// 去除冗余的i
		for len(dq) != 0 && sums[i] <= sums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	if res == 1<<31 {
		return -1
	}
	return res
}

func helper862Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestShortestSubarray(t *testing.T) {
	tests := []struct {
		A      []int
		K      int
		output int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 4, -1},
		{[]int{2, -1, 2}, 3, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, shortestSubarray(ts.A, ts.K))
		})
	}
}
