package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-170/problems/minimum-insertion-steps-to-make-a-string-palindrome/
// ref: https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/discuss/470706/JavaC%2B%2BPython-Longest-Common-Sequence
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i] == s[n-1-j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	return n - dp[n][n]
}

func TestMinInsertions(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"zzazz", 0},
		{"mbadm", 2},
		{"leetcode", 5},
		{"g", 0},
		{"no", 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minInsertions(ts.s))
		})
	}
}
