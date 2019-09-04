package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-partitioning/description/
func partition(s string) [][]string {
	res := [][]string{}
	cur := []string{}
	helper131(&res, cur, s, 0)
	return res
}

func helper131(res *[][]string, cur []string, s string, l int) {
	// 结果
	if len(cur) > 0 && l >= len(s) {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
	}
	// 回溯模板
	for i := l; i < len(s); i++ {
		if helper131Palindrome(s[l : i+1]) {
			if l == i {
				cur = append(cur, string(s[i]))
			} else {
				cur = append(cur, s[l:i+1])
			}
			helper131(res, cur, s, i+1)
			cur = cur[:len(cur)-1]
		}
	}
}

func helper131Palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func TestPartition(t *testing.T) {
	tests := []struct {
		s      string
		output [][]string
	}{
		{"aab", [][]string{[]string{"a", "a", "b"}, []string{"aa", "b"}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, partition(ts.s))
		})
	}
}
