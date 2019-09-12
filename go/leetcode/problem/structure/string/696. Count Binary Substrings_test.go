package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-binary-substrings/description/

// Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

// Substrings that occur multiple times are counted the number of times they occur.

func countBinarySubstrings(s string) int {
	count, count0, count1 := 0, 0, 0
	prev := rune(s[0])
	for _, r := range s {
		if prev == r {
			if r == '0' {
				count0++
			} else {
				count1++
			}
		} else {
			count += helper696(count0, count1)

			if r == '0' {
				count0 = 1
			} else {
				count1 = 1
			}
		}
		prev = r
	}
	return count + helper696(count0, count1)
}

func helper696(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		s      string
		output int
	}{
		{"00110011", 6},
		// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
		// Notice that some of these substrings repeat and are counted the number of times they occur.
		// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
		{"10101", 4},
		// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countBinarySubstrings(ts.s))
		})
	}
}
