package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-break/description/

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

// Note:
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.

// 完全背包：物品数量为无限个
// 求解顺序的完全背包问题时，
// 对物品的迭代应该放在最里层，对背包的迭代放在外层，只有这样才能让物品按一定顺序放入背包中。

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict { // 对物品的迭代应该放在最里层
			wLen := len(word)
			if wLen <= i && strings.Compare(word, s[i-wLen:i]) == 0 {
				dp[i] = dp[i] || dp[i-wLen]
			}
		}
	}

	return dp[n]
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		output   bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		// Explanation: Return true because "leetcode" can be segmented as "leet code".
		{"applepenapple", []string{"apple", "pen"}, true},
		// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
		//              Note that you are allowed to reuse a dictionary word.
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"catsandog", []string{}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, wordBreak(ts.s, ts.wordDict))
		})
	}
}
