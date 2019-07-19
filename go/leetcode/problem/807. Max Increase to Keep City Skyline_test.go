package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-increase-to-keep-city-skyline/

// In a 2 dimensional array grid, each value grid[i][j] represents the height of a building located there. We are allowed to increase the height of any number of buildings, by any amount (the amounts can be different for different buildings). Height 0 is considered to be a building as well.

// At the end, the "skyline" when viewed from all four directions of the grid, i.e. top, bottom, left, and right, must be the same as the skyline of the original grid. A city's skyline is the outer contour of the rectangles formed by all the buildings when viewed from a distance. See the following example.

// What is the maximum total sum that the height of the buildings can be increased?

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// gets the row & col skyline
	row := make([]int, len(grid))
	col := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			row[i] = max(row[i], grid[i][j])
			col[j] = max(col[j], grid[i][j])
		}
	}
	//
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += max(grid[i][j], min(row[i], col[j])) - grid[i][j]
		}
	}
	return res
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}, 35},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxIncreaseKeepingSkyline(ts.grid))
		})
	}
}
