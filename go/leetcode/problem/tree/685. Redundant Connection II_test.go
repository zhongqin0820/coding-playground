package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/redundant-connection-ii/

// In this problem, a rooted tree is a directed graph such that, there is exactly one node (the root) for which all other nodes are descendants of this node, plus every node has exactly one parent, except for the root node which has no parents.
// The given input is a directed graph that started as a rooted tree with N nodes (with distinct values 1, 2, ..., N), with one additional directed edge added. The added edge has two different vertices chosen from 1 to N, and was not an edge that already existed.
// The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v] that represents a directed edge connecting nodes u and v, where u is a parent of child v.
// Return an edge that can be removed so that the resulting graph is a rooted tree of N nodes. If there are multiple answers, return the answer that occurs last in the given 2D-array.

// Note:
// The size of the input 2D-array will be between 3 and 1000.
// Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.
// 并查集
func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	var first, last []int

	for k := range edges {
		p, c := edges[k][0], edges[k][1]
		if parent[c] == 0 {
			parent[c] = p
		} else {
			first = []int{parent[c], c}
			last = edges[k]
			edges[k] = nil
			break
		}
	}

	root := parent
	for i := 0; i <= n; i++ {
		root[i] = i
	}

	rootOf := func(i int) int {
		for i != root[i] {
			i = root[i]
		}
		return i
	}
	for _, edge := range edges {
		if edge == nil {
			continue
		}
		p := edge[0]
		c := edge[1]
		r := rootOf(p)

		if r == c {
			if first == nil {
				return edge
			}
			return first
		}

		root[c] = r
	}

	return last
}

func TestFindRedundantDirectedConnection(t *testing.T) {
	tests := []struct {
		edges  [][]int
		output []int
	}{
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		// Explanation: The given directed graph will be like this:
		//   1
		//  / \
		// v   v
		// 2-->3
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}, []int{4, 1}},
		// Explanation: The given directed graph will be like this:
		// 5 <- 1 -> 2
		//      ^    |
		//      |    v
		//      4 <- 3
		{[][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}, []int{2, 1}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findRedundantDirectedConnection(ts.edges))
		})
	}
}
