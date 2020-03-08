package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/unique-paths-iii/
func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res, empty, s_i, s_j, e_i, e_j := 0, 1, 0, 0, 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 0:
				empty++
			case 1:
				s_i, s_j = i, j
			case 2:
				e_i, e_j = i, j
			}
		}
	}
	helper980(&grid, &res, &empty, s_i, s_j, m, n, s_i, s_j, e_i, e_j)
	return res
}

func helper980(grid *[][]int, res, empty *int, i, j, m, n, s_i, s_j, e_i, e_j int) {
	if i < 0 || i >= m || j < 0 || j >= n || (*grid)[i][j] < 0 {
		return
	}
	if i == e_i && j == e_j {
		if (*empty) == 0 {
			*res = *res + 1
		}
		return
	}
	(*grid)[i][j] = -2
	*empty = *empty - 1
	helper980(grid, res, empty, i-1, j, m, n, s_i, s_j, e_i, e_j)
	helper980(grid, res, empty, i, j+1, m, n, s_i, s_j, e_i, e_j)
	helper980(grid, res, empty, i+1, j, m, n, s_i, s_j, e_i, e_j)
	helper980(grid, res, empty, i, j-1, m, n, s_i, s_j, e_i, e_j)
	(*grid)[i][j] = 0
	*empty = *empty + 1
}

func TestUniquePathsIII(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}, 2},
		{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}, 4},
		{[][]int{{0, 1}, {2, 0}}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, uniquePathsIII(ts.grid))
		})
	}
}
