package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/
func longestSubsequence(arr []int, difference int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	res, n := 1, len(arr)
	m := make(map[int]int, n) // m[arr[i]]i 表示元素arr[i]对应的下标i，从而获得对应的dp[i]
	dp := make([]int, n)      // dp[i] 表示以arr[i]为最后一个元素的最长序列长度
	// dp[i] = dp[m[arr[i]-difference]] + 1

	dp[0] = 1
	m[arr[0]] = 0

	for i := 1; i < n; i++ {
		if idx, ok := m[arr[i]-difference]; ok {
			dp[i] = dp[idx] + 1
			m[arr[i]] = i
		} else if idx, ok = m[arr[i]]; ok && 1 < dp[idx] {
			dp[i] = dp[idx]
		} else {
			dp[i] = 1
			m[arr[i]] = i
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		arr        []int
		difference int
		output     int
	}{
		{[]int{1, 2, 3, 4}, 1, 4},
		{[]int{1, 3, 5, 7}, 1, 1},
		{[]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, longestSubsequence(ts.arr, ts.difference))
		})
	}
}
