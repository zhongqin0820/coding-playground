package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 1 {
		return n
	}
	m := [256]int{}
	for i := range m {
		m[i] = -1
	}
	// i 表示至少第2次出现的字符位置
	i, max := 0, 0
	for j, _ := range s {
		if m[s[j]] >= i {
			// 至少第2次出现更新字符位置
			i = m[s[j]] + 1
		} else if j+1-i > max {
			// 结果更新
			max = j + 1 - i
		}
		m[s[j]] = j
	}
	return max
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"", 0},
		{"a", 1},
		{"abcabcbb", 3},
		// Explanation: The answer is "abc", with the length of 3.
		{"bbbbb", 1},
		// Explanation: The answer is "b", with the length of 1.
		{"pwwkew", 3},
		// Explanation: The answer is "wke", with the length of 3.
		// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, lengthOfLongestSubstring(ts.input))
		})
	}
}
