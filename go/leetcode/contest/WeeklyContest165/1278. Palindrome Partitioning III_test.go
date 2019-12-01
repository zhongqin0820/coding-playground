package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-partitioning-iii/
func palindromePartition(s string, k int) int {
	if len(s) < 1 || len(s) > 100 {
		return -1
	}
	// TODO(zoking)
	return 0
}

func TestPalindromePartition(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output int
	}{
		{"abc", 2, 1},
		{"aabbc", 3, 0},
		{"leetcode", 8, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, palindromePartition(ts.s, ts.k))
		})
	}
}
