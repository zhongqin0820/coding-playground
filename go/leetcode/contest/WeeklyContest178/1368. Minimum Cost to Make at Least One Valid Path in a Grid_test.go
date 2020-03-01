package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ref: https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/discuss/524886/Python-BFS-and-DFS
func minCost(grid [][]int) int {
	// TODO:Zoking:2020-03-01
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	dp := make([][]int, m)
	temp := make([]int, n)
	for i := range temp {
		temp[i] = -1
	}
	for i := range dp {
		dp[i] = make([]int, n)
		copy(dp[i], temp)
	}
	//
	bfs := make([][2]int, 0, 1024)
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || dp[i][j] != -1 {
			return
		}
		dp[i][j] = res
		bfs = append(bfs, [2]int{i, j})
		dfs(i+dir[grid[i][j]-1][0], j+dir[grid[i][j]-1][1])
	}
	dfs(0, 0)
	for len(bfs) != 0 {
		res += 1
		bfs2 := make([][2]int, len(bfs))
		copy(bfs2, bfs)
		bfs = [][2]int{}
		for _, dir_ := range dir {
			for _, bfs_ := range bfs2 {
				dfs(dir_[0]+bfs_[0], dir_[1]+bfs_[1])
			}
		}
	}
	return dp[m-1][n-1]
}

func TestMinCost(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}, 3},
		{[][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}, 0},
		{[][]int{{1, 2}, {4, 3}}, 1},
		{[][]int{{2, 2, 2}, {2, 2, 2}}, 3},
		{[][]int{{4}}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minCost(ts.grid))
		})
	}
}
