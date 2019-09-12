package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow/

// Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.
// Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.
// Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

// Note:
// 1. The order of returned grid coordinates does not matter.
// 1. Both m and n are less than 150.

func pacificAtlanticBFS(mat [][]int) [][]int {
	res := [][]int{}
	if len(mat) == 0 || len(mat[0]) == 0 {
		return res
	}

	m, n := len(mat), len(mat[0])

	// p[i][j] 表示，[i][j] 可以让水流到 Pacific 的点
	// a[i][j] 表示，[i][j] 可以让水流到 Atlantic 的点
	p, a := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}
	// pQueue 是所有能够让水流到 Pacific 的点的队列
	// aQueue 是所有能够让水流到 Atlantic 的点的队列
	// 初始化 pQueue 和 aQueue
	pQueue := [][]int{}
	aQueue := [][]int{}
	// 左边可进pQueue，右边进aQueue
	for i := 0; i < m; i++ {
		p[i][0] = true
		pQueue = append(pQueue, []int{i, 0})
		a[i][n-1] = true
		aQueue = append(aQueue, []int{i, n - 1})
	}
	// 上边进pQueue，下边进aQueue
	for j := 0; j < n; j++ {
		p[0][j] = true
		pQueue = append(pQueue, []int{0, j})
		a[m-1][j] = true
		aQueue = append(aQueue, []int{m - 1, j})
	}

	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	bfs := func(queue [][]int, rec [][]bool) {
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			for _, d := range ds {
				i, j := c[0]+d[0], c[1]+d[1]
				if 0 <= i && i < m &&
					0 <= j && j < n &&
					!rec[i][j] &&
					mat[c[0]][c[1]] <= mat[i][j] { // i,j可达是流动方向则进队列且
					rec[i][j] = true
					queue = append(queue, []int{i, j})
				}
			}
		}
	}

	bfs(pQueue, p)
	bfs(aQueue, a)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output [][]int
	}{
		{
			[][]int{},
			[][]int{},
		},

		{
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, pacificAtlanticBFS(ts.matrix))
		})
	}
}
