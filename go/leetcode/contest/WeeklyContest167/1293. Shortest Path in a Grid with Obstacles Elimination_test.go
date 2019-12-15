package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-167/problems/shortest-path-in-a-grid-with-obstacles-elimination/
func shortestPath(grid [][]int, k int) int {
	// TODO(zoking)[2019-12-15]: https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/discuss/451782/Java-Solution-Using-DFS-and-Memoization
	return -1
}

func TestShortestPath(t *testing.T) {
	tests := []struct {
		grid   [][]int
		k      int
		output int
	}{
		{
			[][]int{
				[]int{0, 0, 0},
				[]int{1, 1, 0},
				[]int{0, 0, 0},
				[]int{0, 1, 1},
				[]int{0, 0, 0},
			},
			1,
			6,
		},
		{
			[][]int{
				[]int{0, 1, 1},
				[]int{1, 1, 1},
				[]int{1, 0, 0},
			},
			1,
			-1,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, shortestPath(ts.grid, ts.k))
		})
	}
}
