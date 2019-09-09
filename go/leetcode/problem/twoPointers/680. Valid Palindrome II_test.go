package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-palindrome-ii/description/
func validPalindrome(s string) bool {
	return helper680([]byte(s), 0, len(s)-1, false)
}

func helper680(bs []byte, i, j int, flag bool) bool {
	for i < j {
		if bs[i] != bs[j] {
			if flag {
				return false
			}
			return helper680(bs, i+1, j, true) ||
				helper680(bs, i, j-1, true)
		}
		i++
		j--
	}
	return true
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{
		{"aba", true},
		{"abca", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, validPalindrome(ts.s))
		})
	}
}
