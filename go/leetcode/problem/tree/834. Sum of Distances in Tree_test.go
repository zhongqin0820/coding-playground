package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-distances-in-tree/

// An undirected, connected tree with N nodes labelled 0...N-1 and N-1 edges are given.
// The ith edge connects nodes edges[i][0] and edges[i][1] together.
// Return a list ans, where ans[i] is the sum of the distances between node i and all other nodes.

func sumOfDistancesInTree(N int, edges [][]int) []int {
	t := make([][]int, N)
	ar1, ar2 := make([]int, N), make([]int, N)
	for _, e := range edges {
		a, b := e[0], e[1]
		t[a] = append(t[a], b)
		t[b] = append(t[b], a)
	}

	var cal1 func(int, int)
	cal1 = func(c, a int) {
		ar2[a] = 1
		for _, b := range t[a] {
			if b != c {
				cal1(a, b)
				ar2[a] += ar2[b]
				ar1[a] += ar1[b] + ar2[b]
			}
		}
	}

	var cal2 func(int, int)
	cal2 = func(c, a int) {
		ar1[a] = ar1[c] + N - (ar2[a] << 1)
		for _, b := range t[a] {
			if b != c {
				cal2(a, b)
			}
		}
	}

	cal1(0, 0)
	for _, b := range t[0] {
		cal2(0, b)
	}
	return ar1
}

func TestSumOfDistancesInTree(t *testing.T) {
	tests := []struct {
		N      int
		edges  [][]int
		output []int
	}{
		{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sumOfDistancesInTree(ts.N, ts.edges))
		})
	}
}
