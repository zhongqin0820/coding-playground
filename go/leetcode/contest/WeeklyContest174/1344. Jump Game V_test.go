package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ref: https://leetcode.com/problems/jump-game-v/discuss/496520/Top-Down-DP-O(nd)
func maxJumps(arr []int, d int) int {
	var dp = make([]int, 1001)
	res := 0
	// 遍历每一个idx
	for i := 0; i < len(arr); i++ {
		res = max(res, dfs(arr, dp, i, d))
	}
	return res
}

func dfs(arr, dp []int, i, d int) int {
	if dp[i] != 0 {
		return dp[i]
	}
	res := 1
	// (i, i+d]
	for j := i + 1; j <= min(i+d, len(arr)-1) && arr[j] < arr[i]; j++ {
		res = max(res, 1+dfs(arr, dp, j, d))
	}
	// [i-d, i)
	for j := i - 1; j >= max(0, i-d) && arr[j] < arr[i]; j-- {
		res = max(res, 1+dfs(arr, dp, j, d))
	}
	dp[i] = res
	return dp[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func TestMaxJumps(t *testing.T) {
	tests := []struct {
		arr    []int
		d      int
		output int
	}{
		{[]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2, 4},
		{[]int{3, 3, 3, 3, 3}, 3, 1},
		{[]int{7, 6, 5, 4, 3, 2, 1}, 1, 7},
		{[]int{7, 1, 7, 1, 7, 1}, 2, 2},
		{[]int{66}, 1, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxJumps(ts.arr, ts.d))
		})
	}
}
