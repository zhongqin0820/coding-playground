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
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, d := range directions {
			dfs(i+d[0], j+d[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '0' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}

func numIslandsBFS(grid [][]byte) int {
	// 获取 grid 的大小
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// bfs 搜索时，存放点坐标的队列
	x := make([]int, 0, m*n)
	y := make([]int, 0, m*n)

	// 往队列中添加 (i,j) 点，并修改 (i,j) 点的值
	// 避免重复搜索
	var add = func(i, j int) {
		x = append(x, i)
		y = append(y, j)
		grid[i][j] = '0'
	}

	// 从坐标队列中，取出坐标点
	var pop = func() (int, int) {
		i := x[0]
		x = x[1:]
		j := y[0]
		y = y[1:]
		return i, j
	}

	var bfs = func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}
		add(i, j)

		for len(x) > 0 {
			i, j = pop()
			// 搜索 (i,j) 点的 上下左右 四个方位
			if 0 <= i-1 && grid[i-1][j] == '1' {
				add(i-1, j)
			}
			if 0 <= j-1 && grid[i][j-1] == '1' {
				add(i, j-1)
			}
			if i+1 < m && grid[i+1][j] == '1' {
				add(i+1, j)
			}
			if j+1 < n && grid[i][j+1] == '1' {
				add(i, j+1)
			}
		}
		return 1
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += bfs(i, j)
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
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numIslands(ts.grid))
			// ast.Equal(ts.output, numIslandsBFS(ts.grid))
		})
	}
}
