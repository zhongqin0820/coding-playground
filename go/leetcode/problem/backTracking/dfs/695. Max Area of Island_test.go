package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-area-of-island/submissions/

// Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
// Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
// Note: The length of each dimension in the given grid does not exceed 50.

func maxAreaOfIsland(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxArea := 0
	for i := range grid {
		for j := range grid[i] {
			maxArea = max(maxArea, helper695(grid, i, j))
		}
	}
	return maxArea
}

// 返回与 [i,j] 处联通的岛的面积
func helper695(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	area := 1
	if i != 0 {
		area += helper695(grid, i-1, j)
	}
	if j != 0 {
		area += helper695(grid, i, j-1)
	}
	if i != len(grid)-1 {
		area += helper695(grid, i+1, j)
	}
	if j != len(grid[0])-1 {
		area += helper695(grid, i, j+1)
	}
	return area
}

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}, 4},
		{
			[][]int{
				{1},
			},
			1,
		},

		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			0,
			// Given the above grid, return 0.
		},

		{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
			// Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxAreaOfIsland(ts.grid))
		})
	}
}
