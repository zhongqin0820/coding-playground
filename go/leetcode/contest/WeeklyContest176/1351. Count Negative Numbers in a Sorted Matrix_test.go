package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
func countNegatives(grid [][]int) int {
	res := 0
	for i := 0; grid != nil && i < len(grid); i++ {
		for j := 0; grid[i] != nil && j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				res++
			}
		}
	}
	return res
}

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
		{[][]int{{3, 2}, {1, 0}}, 0},
		{[][]int{{1, -1}, {-1, -1}}, 3},
		{[][]int{{-1}}, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countNegatives(ts.grid))
		})
	}
}
