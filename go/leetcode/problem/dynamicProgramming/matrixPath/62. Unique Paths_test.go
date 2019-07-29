package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/unique-paths/description/

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?

// Note: m and n will be at most 100.

func uniquePaths(m int, n int) int {
	// init
	dp := [100][100]int{}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// dp[i][j]=dp[i][j-1]+dp[i-1][j]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsFormula(m int, n int) int { // C_m^n
	s := m + n - 2 //total moves
	d := m - 1     //downward moves
	res := 1
	for i := 1; i <= d; i++ {
		res = res * (s - d + i) / i
	}
	return res
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n   int
		output int
	}{
		{3, 2, 3},
		// Explanation:
		// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
		// 1. Right -> Right -> Down
		// 2. Right -> Down -> Right
		// 3. Down -> Right -> Right
		{7, 3, 28},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, uniquePaths(ts.m, ts.n))
			ast.Equal(ts.output, uniquePathsFormula(ts.m, ts.n))
		})
	}
}
