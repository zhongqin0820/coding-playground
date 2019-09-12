package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/distinct-subsequences/

// Given a string S and a string T, count the number of distinct subsequences of S which equals T.

// A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = 1
	}

	var prev int
	for j := 0; j < n; j++ {
		dp[j], prev = 0, dp[j]

		for i := j + 1; i <= m; i++ {
			if t[j] == s[i-1] {
				dp[i], prev = dp[i-1]+prev, dp[i]
			} else {
				dp[i], prev = dp[i-1], dp[i]
			}
		}
	}
	return dp[m]
}

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s, t   string
		output int
	}{
		{"rabbbit", "rabbit", 3},
		// Explanation:
		// As shown below, there are 3 ways you can generate "rabbit" from S.
		// (The caret symbol ^ means the chosen letters)

		// rabbbit
		// ^^^^ ^^
		// rabbbit
		// ^^ ^^^^
		// rabbbit
		// ^^^ ^^^
		{"babgbag", "bag", 5},
		// Explanation:
		// As shown below, there are 5 ways you can generate "bag" from S.
		// (The caret symbol ^ means the chosen letters)

		// babgbag
		// ^^ ^
		// babgbag
		// ^^    ^
		// babgbag
		// ^    ^^
		// babgbag
		//   ^  ^^
		// babgbag
		//     ^^^
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numDistinct(ts.s, ts.t))
		})
	}
}
