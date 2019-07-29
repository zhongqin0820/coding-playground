package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/ones-and-zeroes/description/

// In the computer world, use restricted resource you have to generate maximum benefit is what we always want to pursue.
// For now, suppose you are a dominator of m 0s and n 1s respectively. On the other hand, there is an array with strings consisting of only 0s and 1s.
// Now your task is to find the maximum number of strings that you can form with given m 0s and n 1s. Each 0 and 1 can be used at most once.

// Note:
// The given numbers of 0s and 1s will both not exceed 100
// The size of given string array won't exceed 600.

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 为 i 个 '0' 和 j 个 '1' 所能形成的最多的 s 的个数
	dp := make([][]int, m+1) //m个0
	for i := range dp {
		dp[i] = make([]int, n+1) //n个1
	}

	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for i := m; i-zeros >= 0; i-- {
			for j := n; j-ones >= 0; j-- {
				// 更新 dp[i][j]
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		strs   []string
		m      int
		n      int
		output int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		// Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s, which are “10,”0001”,”1”,”0”
		{[]string{"10", "0", "1"}, 1, 1, 2},
		// Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMaxForm(ts.strs, ts.m, ts.n))
		})
	}
}
