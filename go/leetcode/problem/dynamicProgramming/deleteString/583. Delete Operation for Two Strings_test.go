package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/delete-operation-for-two-strings/description/

// Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.

// Note:
// The length of given words won't exceed 500.
// Characters in given words can only be lower-case letters.

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	// dp[i][j] == k 表示 word1[:i] 和 word2[:j] 的最大公共子序列的长度为 k
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return n1 + n2 - dp[n1][n2]*2
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1  string
		word2  string
		output int
	}{
		{"sea", "eat", 2},
		// Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minDistance(ts.word1, ts.word2))
		})
	}
}
