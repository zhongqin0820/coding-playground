package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shortest-palindrome/

// Given a string s, you are allowed to convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	// 找到 (s+"#"+reverse(s)) 的最大匹配的下标
	i := helper214(s + "#" + helper214Reverse(s))
	s = helper214Reverse(s[i:]) + s
	return s
}

func helper214(s string) int {
	n := len(s)
	table := make([]int, n)
	klen, i := 0, 1
	for i < n {
		if s[i] == s[klen] {
			klen++
			table[i] = klen
			i++
		} else {
			if klen == 0 {
				table[i] = 0
				i++
			} else {
				klen = table[klen-1]
			}
		}
	}
	//
	i = table[n-1]
	return i
}

func helper214Reverse(s string) string {
	n := len(s)
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = s[n-i-1]
	}
	return string(bs)
}

func TestShortestPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"aacecaaa", "aaacecaaa"},
		{"abcd", "dcbabcd"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, shortestPalindrome(ts.s))
		})
	}
}
