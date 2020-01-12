package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-171/problems/number-of-operations-to-make-network-connected/
// ref: https://leetcode.com/problems/number-of-operations-to-make-network-connected/discuss/477676/Union-Find%3A-Count-groups-and-extra-cables
func makeConnected(n int, connections [][]int) int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}
	// count extra connections
	extra_cables := 0
	for _, c := range connections {
		i, j := helper1319(parent, c[0]), helper1319(parent, c[1])
		if i == j {
			extra_cables++
		} else {
			parent[i] += parent[j]
			parent[j] = i
		}
	}
	// count the isolated node
	groups := 0
	for _, p := range parent {
		if p < 0 {
			groups++
		}
	}
	// return result
	if extra_cables >= groups-1 {
		return groups - 1
	}
	return -1
}

// find node i-th parent
func helper1319(parent []int, i int) int {
	if parent[i] < 0 {
		return i
	}
	parent[i] = helper1319(parent, parent[i])
	return parent[i]
}

func TestMakeConnected(t *testing.T) {
	// Constraints:
	//     1 <= n <= 10^5
	//     1 <= connections.length <= min(n*(n-1)/2, 10^5)
	//     connections[i].length == 2
	//     0 <= connections[i][0], connections[i][1] < n
	//     connections[i][0] != connections[i][1]
	//     There are no repeated connections.
	//     No two computers are connected by more than one cable.
	tests := []struct {
		n           int
		connections [][]int
		output      int
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
		{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, makeConnected(ts.n, ts.connections))
		})
	}
}
