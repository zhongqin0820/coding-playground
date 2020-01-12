package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// find the initial island
	q := make([][]int, 0, 100)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// count the step
	steps := 0
	for l := len(q); l > 0; l = len(q) {
		q1 := make([][]int, 0, 100)
		steps++
		for i := 0; i < l; i++ {
			p := q[i]
			for _, dir := range dirs {
				p_ := []int{p[0] + dir[0], p[1] + dir[1]}
				if p_[0] < 0 || p_[0] >= m || p_[1] < 0 || p_[1] >= n || grid[p_[0]][p_[1]] != 0 {
					continue
				}
				grid[p_[0]][p_[1]] = steps
				q1 = append(q1, p_)
			}
		}
		q = q1
	}
	if steps == 1 {
		return -1
	}
	return steps - 1
}

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			2,
		},
		{
			[][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			4,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxDistance(ts.grid))
		})
	}
}
