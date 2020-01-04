package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pos struct {
	i, j int
}

func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	fresh := 0
	res := 0
	q := make([]pos, 0, 100)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				q = append(q, pos{i: i, j: j})
			}
		}
	}
	if fresh == 0 {
		return res
	}
	res--
	dirs := []pos{pos{-1, 0}, pos{0, -1}, pos{1, 0}, pos{0, 1}}
	for l := len(q); l > 0; l = len(q) {
		for l_ := 0; l_ < l; l_++ {
			p := q[l_]
			for _, dir := range dirs {
				t := pos{p.i + dir.i, p.j + dir.j}
				if t.i < 0 || t.j < 0 || t.i >= n || t.j >= m || grid[t.i][t.j] != 1 {
					continue
				}
				grid[t.i][t.j] = 2
				fresh--
				q = append(q, t)
			}
		}
		q = q[l:]
		res++
	}
	if fresh > 0 {
		return -1
	}
	return res
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{
			[][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}},
			4,
		},
		{
			[][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}},
			-1,
		},
		{
			[][]int{[]int{0, 2}},
			0,
		},
		{
			[][]int{{0}},
			0,
		},
		{
			[][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}},
			1,
		},
		{
			[][]int{{1}},
			-1,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, orangesRotting(ts.grid))
		})
	}
}
