package problem

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-characters-by-frequency/description/
func frequencySort(s string) string {
	// 统计字频
	r := [256]int{}
	for i := range s {
		r[s[i]]++
	}
	// 根据字母与频率生成一个string，将所有string放到一个切片中
	ss := make([]string, 0, len(s))
	for i := range r {
		if r[i] == 0 {
			continue
		}
		ss = append(ss, strings.Repeat(string(byte(i)), r[i]))
	}
	sort.Sort(segments(ss))
	res := ""
	for _, s := range ss {
		res += s
	}
	return res
}

type segments []string

func (s segments) Len() int      { return len(s) }
func (s segments) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s segments) Less(i, j int) bool {
	// 长度相同按照字母序排序
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	// 否则长度优先
	return len(s[i]) > len(s[j])
}

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"tree", "eert"},
		{"cccaaa", "aaaccc"},
		{"Aabb", "bbAa"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, frequencySort(ts.s))
		})
	}
}
