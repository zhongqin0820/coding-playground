package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-mountain-array/
func validMountainArray(A []int) bool {
	if A == nil || len(A) < 3 {
		return false
	}
	n := len(A)
	l, r := 1, n-2
	for ; l < n; l++ {
		if A[l] <= A[l-1] {
			break
		}
	}
	for ; r >= 0; r-- {
		if A[r] <= A[r+1] {
			break
		}
	}
	return l-1 == r+1 && l-1 != 0 && r+1 != n-1
}

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		A      []int
		output bool
	}{
		{[]int{0, 2, 3, 4, 5, 2, 1, 0}, true},
		{[]int{0, 2, 3, 3, 5, 2, 1, 0}, false},
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			false,
		},

		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			false,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, validMountainArray(ts.A))
		})
	}
}
