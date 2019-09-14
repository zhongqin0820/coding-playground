package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/power-of-three/description/
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		n      int
		output bool
	}{
		{-1, false},
		{8, false},
		{2, false},
		{3, true},
		{81, true},
		{243, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPowerOfThree(ts.n))
		})
	}
}
