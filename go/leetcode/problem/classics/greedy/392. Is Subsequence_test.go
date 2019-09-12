package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	if n > m {
		return false
	}
	i, j := 0, 0
	for ; i < n && j < m; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == n
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		output bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isSubsequence(ts.s, ts.t))
		})
	}
}
