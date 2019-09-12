package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/redundant-connection/description/
func findRedundantConnection(edges [][]int) []int {
	n, parent := len(edges), make([]int, len(edges)+1)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var i int
	var e []int
	for i, e = range edges {
		// 一条边起终点对应的父亲节点
		f, t := e[0], e[1]
		pf := helper684(parent, f)
		pt := helper684(parent, t)
		// 联通
		if pf == pt {
			break
		}
		parent[pf] = pt
	}
	return edges[i]
}

func helper684(parent []int, f int) int {
	for f != parent[f] {
		f = parent[f]
	}
	return f
}

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		edges  [][]int
		output []int
	}{
		{
			[][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}},
			[]int{4, 10},
		},
		{
			[][]int{{3, 7}, {1, 4}, {2, 8}, {1, 6}, {7, 9}, {6, 10}, {1, 7}, {2, 3}, {8, 9}, {5, 9}},
			[]int{8, 9},
		},
		{
			[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}},
			[]int{1, 5},
		},
		{
			[][]int{{1, 3}, {3, 4}, {1, 5}, {3, 5}, {2, 3}},
			[]int{3, 5},
		},
		{
			[][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}},
			[]int{1, 3},
		},
		{
			[][]int{{1, 2}, {1, 3}, {2, 3}},
			[]int{2, 3},
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			[]int{1, 4},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findRedundantConnection(ts.edges))
		})
	}
}
