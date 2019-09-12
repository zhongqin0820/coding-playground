package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/perfect-squares/description/

// Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	// dp[i] 表示 the least number of perfect square numbers which sum to i
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i <= n; i++ {
			if dp[i] > dp[i-p]+1 {
				// i = ( i - p ) + p，p 是 平方数
				// dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}
	// log.Println(dp)
	return dp[n]
}

func TestNumSquares(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{12, 3},
		// Explanation: 12 = 4 + 4 + 4.
		{13, 2},
		// Explanation: 13 = 4 + 9.
		{1, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numSquares(ts.n))
		})
	}
}
