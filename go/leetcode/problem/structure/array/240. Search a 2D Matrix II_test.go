package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/search-a-2d-matrix-ii/description/

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
		{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, searchMatrix(ts.matrix, ts.target))
		})
	}
}
