package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/
func findLongestWord(s string, d []string) string {
	// 排序后的符合要求
	sort.Sort(stringSlice(d))
	// 判断是否是子串
	for i := range d {
		if isSub(s, d[i]) {
			return d[i]
		}
	}
	return ""
}

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(i, j int) bool {
	// 长度相等，字典序
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	// 否则长度大的在前
	return len(s[i]) > len(s[j])
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isSub(s, sub string) bool {
	if len(s) < len(sub) {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == len(sub)
}

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		s      string
		d      []string
		output string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findLongestWord(ts.s, ts.d))
		})
	}
}
