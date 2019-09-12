package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decode-ways/description/

// A message containing letters from A-Z is being encoded to numbers using the following mapping:

//     'A' -> 1
//     'B' -> 2
//     ...
//     'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.

func numDecodings(s string) int {
	res := 1
	pre := 1
	for i := 0; i < len(s); i++ {
		cur := 0
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7')) {
			cur += pre
		}
		if s[i] > '0' {
			cur += res
		}
		pre, res = res, cur
	}
	return res
}

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"12", 2},
		// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
		{"226", 3},
		// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numDecodings(ts.s))
		})
	}
}
