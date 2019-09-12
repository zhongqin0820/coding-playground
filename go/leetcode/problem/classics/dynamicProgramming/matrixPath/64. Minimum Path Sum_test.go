package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-path-sum/description/

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

// Note: You can only move either down or right at any point in time.

func minPathSum(grid [][]int) int {
	// edge check
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// init
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	// init
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	//
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = grid[i][j] + helper64(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[n-1][m-1]
}

func helper64(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		// Explanation: Because the path 1→3→1→1→1 minimizes the sum.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minPathSum(ts.grid))
		})
	}
}
