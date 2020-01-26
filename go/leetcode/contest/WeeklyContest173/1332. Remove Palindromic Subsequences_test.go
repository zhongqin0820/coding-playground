package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removePalindromeSub(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return 2
		}
	}
	return 1
}

func TestRemovePalindromeSub(t *testing.T) {
	// Constraints:
	//     0 <= s.length <= 1000
	//     s only consists of letters 'a' and 'b'
	tests := []struct {
		s      string
		output int
	}{
		{"ababa", 1},
		{"abb", 2},
		{"baabb", 2},
		{"", 0},
		{"aabb", 2},
		{"bbaabaaa", 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, removePalindromeSub(ts.s))
		})
	}
}
