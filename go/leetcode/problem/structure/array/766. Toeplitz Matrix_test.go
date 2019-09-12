package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/toeplitz-matrix/description/

// A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.
// Now given an M x N matrix, return True if and only if the matrix is Toeplitz.

func isToeplitzMatrix(matrix [][]int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return true
	}
	n, m := len(matrix), len(matrix[0])
	for i := 0; i+1 < n; i++ {
		for j := 0; j+1 < m; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

func TestIsToeplitzMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output bool
	}{
		{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isToeplitzMatrix(ts.matrix))
		})
	}
}
