package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-islands/description/
//
// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	dx, dy := [4]int{-1, 1, 0, 0}, [4]int{0, 0, -1, 1}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				// bfs one island and marked the nodes
				q := make([][2]int, 1, n*m)
				q[0] = [2]int{i, j}
				grid[i][j] = '0'
				for idx := 0; idx < len(q); idx++ {
					node := q[idx]
					for k := 0; k < 4; k++ {
						i_ := node[0] + dx[k]
						j_ := node[1] + dy[k]
						if i_ >= 0 && j_ >= 0 && i_ < n && j_ < m && grid[i_][j_] == '1' {
							q = append(q, [2]int{i_, j_})
							grid[i_][j_] = '0'
						}
					}
				}
				// add one island
				res++
			}
		}
	}
	return res
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		output int
	}{
		{
			[][]byte{},
			0,
		},

		{
			[][]byte{
				[]byte("111001"),
				[]byte("010001"),
				[]byte("111111"),
			},
			1,
		},

		{
			[][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			1,
		},

		{
			[][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			3,
		},

		{
			[][]byte{
				[]byte("11111011111111101011"),
				[]byte("01111111111110111110"),
				[]byte("10111001101111111111"),
				[]byte("11110111111111111111"),
				[]byte("10011111111111111111"),
				[]byte("10111111011101110111"),
				[]byte("01111111111101101111"),
				[]byte("11111111111101111011"),
				[]byte("11111111110111111111"),
				[]byte("11111111111111111111"),
				[]byte("01111111011111111111"),
				[]byte("01111111011111111111"),
				[]byte("11111111111111111111"),
				[]byte("11111111111111111111"),
				[]byte("11111011111110111111"),
				[]byte("10111110111011110111"),
				[]byte("11111111111101111110"),
				[]byte("11111111111110111100"),
				[]byte("11111111111111111111"),
				[]byte("11111111111111111111"),
				[]byte("11111111111111111111"),
			},
			1,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numIslands(ts.grid))
		})
	}
}
