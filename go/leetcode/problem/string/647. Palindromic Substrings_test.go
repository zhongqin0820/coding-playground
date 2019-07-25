package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindromic-substrings/description/

// Given a string, your task is to count how many palindromic substrings in this string.

// The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// 以i为中心
		count += helper647(s, i, i)
		// 以i,i+1为中心
		count += helper647(s, i, i+1)
	}
	return count
}

// 向两边扩展
func helper647(s string, i, j int) int {
	res := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res++
		i--
		j++
	}
	return res
}

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"abc", 3},
		{"aaa", 6},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countSubstrings(ts.s))
		})
	}
}
