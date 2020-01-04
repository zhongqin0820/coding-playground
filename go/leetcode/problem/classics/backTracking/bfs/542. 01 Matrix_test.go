package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/01-matrix/

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	MIN := m * n

	// 根据 左方 和 上方 的格子，更新 (i,j)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			matrix[i][j] = MIN

			if 0 <= i-1 {
				matrix[i][j] = helper542min(matrix[i][j], matrix[i-1][j]+1)
			}
			if 0 <= j-1 {
				matrix[i][j] = helper542min(matrix[i][j], matrix[i][j-1]+1)
			}
		}
	}

	// 根据 右方 和 下方 的格子，更新 (i,j)
	for i := m - 1; 0 <= i; i-- {
		for j := n - 1; 0 <= j; j-- {
			if matrix[i][j] == 0 {
				continue
			}

			if i+1 < m {
				matrix[i][j] = helper542min(matrix[i][j], matrix[i+1][j]+1)
			}
			if j+1 < n {
				matrix[i][j] = helper542min(matrix[i][j], matrix[i][j+1]+1)
			}
		}
	}

	return matrix
}

func helper542min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, updateMatrix(ts.matrix))
		})
	}
}
