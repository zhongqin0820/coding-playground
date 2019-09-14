package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-perfect-square/description/
func isPerfectSquare(num int) bool {
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	return r*r == num
}

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num    int
		output bool
	}{
		{16, true},
		{14, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isPerfectSquare(ts.num))
		})
	}
}
