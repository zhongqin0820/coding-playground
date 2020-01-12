package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/repeated-string-match/
func repeatedStringMatch(A string, B string) int {
	n, m := len(A), len(B)
	prefTable := make([]int, m+1)
	for sp, pp := 1, 0; sp < m; {
		if B[pp] == B[sp] {
			sp++
			pp++
			prefTable[sp] = pp
		} else if pp == 0 {
			sp++
			prefTable[sp] = pp
		} else {
			pp = prefTable[pp]
		}
	}

	for i, j := 0, 0; i < n; i, j = i+helper686max(1, j-prefTable[j]), prefTable[j] {
		for j < m && A[(i+j)%n] == B[j] {
			j++
		}
		if j == m {
			if (i+j)%n == 0 {
				return (i + j) / n
			} else {
				return (i+j)/n + 1
			}
		}
	}
	return -1
}

func helper686max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestRepeatedStringMatch(t *testing.T) {
	tests := []struct {
		A      string
		B      string
		output int
	}{
		{"abcd", "cdabcdab", 3},
		{"abc", "cabcabca", 4},
		{
			"ba",
			"ab",
			2,
		},
		{
			"abc",
			"aabcbabcc",
			-1,
		},
		{
			"abcd",
			"cdabdab",
			-1,
		},
		{
			"aa",
			"b",
			-1,
		},
		{
			"aa",
			"a",
			1,
		},
		{
			"abcd",
			"bcabcdbc",
			-1,
		},
		{
			"abcd",
			"abcdb",
			-1,
		},
		{
			"abababaaba",
			"aabaaba",
			2,
		},
		{
			"bb",
			"bbbbbbb",
			4,
		},
		{
			"a",
			"aa",
			2,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, repeatedStringMatch(ts.A, ts.B))
		})
	}
}
