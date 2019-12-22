package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-168/problems/maximum-number-of-occurrences-of-a-substring/
// ref: https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/discuss/457577/C%2B%2B-Sliding-window-O(n)
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	res, n := 0, len(s)
	m := make(map[string]int, 2048)
	for l := 0; l < n-minSize+1; l++ {
		sub := s[l : l+minSize]
		b := make(map[rune]struct{}, len(sub))
		for _, c := range sub {
			b[c] = struct{}{}
		}
		if len(b) > maxLetters {
			continue
		}
		m[sub]++
		if res < m[sub] {
			res = m[sub]
		}
	}
	return res
}

func TestMaxFreq(t *testing.T) {
	tests := []struct {
		s          string
		maxLetters int
		minSize    int
		maxSize    int
		output     int
	}{
		{"aababcaab", 2, 3, 4, 2},
		{"aaaa", 1, 3, 3, 2},
		{"aabcabcab", 2, 2, 3, 3},
		{"abcde", 2, 3, 3, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxFreq(ts.s, ts.maxLetters, ts.minSize, ts.maxSize))
		})
	}
}
