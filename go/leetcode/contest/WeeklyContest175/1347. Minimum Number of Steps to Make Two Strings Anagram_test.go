package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minSteps(s string, t string) int {
	m := make(map[byte]int, len(s))
	for i := range s {
		m[s[i]]++
		m[t[i]]--
	}
	res := 0
	for _, v := range m {
		if v < 0 {
			res -= v
		}
	}
	return res
}

func TestMinSteps(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		output int
	}{
		{"bab", "aba", 1},
		{"leetcode", "practice", 5},
		{"anagram", "mangaar", 0},
		{"xxyyzz", "xxyyzz", 0},
		{"friend", "family", 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minSteps(ts.s, ts.t))
		})
	}
}
