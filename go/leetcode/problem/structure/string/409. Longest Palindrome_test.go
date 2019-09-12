package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-palindrome/

// Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
// This is case sensitive, for example "Aa" is not considered a palindrome here.

// Note:
// Assume the length of given string will not exceed 1,010.

func longestPalindrome409(s string) int {
	res, odd := 0, 0
	a := [123]int{} //byte('z')=122
	for i := range s {
		a[s[i]]++
	}

	// 统计偶数个元素个数，如果是奇数则减1变偶数，odd=1等价于有一个奇数的字母放在中间
	for i := range a {
		if a[i] == 0 {
			continue
		}

		if a[i]&1 == 0 {
			res += a[i]
		} else {
			res += a[i] - 1
			odd = 1
		}
	}
	return res + odd
}

func TestLongestPalindrome409(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"abccccdd", 7},
		// Explanation:
		// One longest palindrome that can be built is "dccaccd", whose length is 7.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, longestPalindrome409(ts.s))
		})
	}
}
