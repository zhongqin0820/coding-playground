package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	n := len(word)
	if n == 1 {
		return true
	}
	for i := 2; i < n; i++ {
		if (lower(word[i-1]) && !lower(word[i])) ||
			(!lower(word[i-1]) && lower(word[i])) {
			return false
		}
	}
	return (lower(word[0]) && lower(word[1])) || !lower(word[0])
}

func lower(a byte) bool {
	return a-'a' >= 0 && a-'a' <= 26
}

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		word   string
		output bool
	}{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
		{"mL", false},
		{"m", true},
		{"M", true},
		{"GooAle", false},
		{"mRZ", false},
		{"mRz", false},
		{"mrz", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, detectCapitalUse(ts.word))
		})
	}
}
