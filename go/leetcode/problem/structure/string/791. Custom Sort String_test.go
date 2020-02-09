package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/custom-sort-string/
func customSortString(S string, T string) string {
	bT := []byte(T)
	m := make(map[byte]int, len(S))
	for i := range S {
		m[S[i]] = i
	}
	sort.Slice(bT, func(i, j int) bool {
		// 如果都在S中，比较idx
		idx_i, ok_i := m[bT[i]]
		idx_j, ok_j := m[bT[j]]
		if ok_i && ok_j {
			return idx_i < idx_j
		}
		// 否则，按照字典序
		if !ok_i && !ok_j {
			return bT[i] < bT[j]
		}
		// 否则出现在S中的优先
		return ok_i
	})
	return string(bT)
}

func TestCustomSortString(t *testing.T) {
	tests := []struct {
		S      string
		T      string
		output string
	}{
		{"cba", "abcd", "cbad"},
		{"disqyr", "iwyrysqrdj", "disqyyrrjw"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, customSortString(ts.S, ts.T))
		})
	}
}
