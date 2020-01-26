package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// TODO[zoking](2020-01-26): ref: https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/discuss/490312/Python-Floyd
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = 1 << 31
		}
	}
	for _, edge := range edges {
		dis[edge[0]][edge[1]] = edge[2]
		dis[edge[1]][edge[0]] = edge[2]
	}
	for i := range dis {
		dis[i][i] = 0
	}
	// Floyd
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][j] > dis[i][k]+dis[k][j] {
					dis[i][j] = dis[i][k] + dis[k][j]
				}
			}
		}
	}
	// check
	res := 1 << 31
	j := 0
	for i := 0; i < n; i++ {
		temp := 0
		for _, d := range dis[i] {
			if d <= distanceThreshold {
				temp += d
			}
		}

		if temp < res {
			res = temp
			j = i
		}
	}
	return j
}

func TestFindTheCity(t *testing.T) {
	// Constraints:
	//     2 <= n <= 100
	//     1 <= edges.length <= n * (n - 1) / 2
	//     edges[i].length == 3
	//     0 <= fromi < toi < n
	//     1 <= weighti, distanceThreshold <= 10^4
	//     All pairs (fromi, toi) are distinct.

	tests := []struct {
		n                 int
		edges             [][]int
		distanceThreshold int
		output            int
	}{
		{4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4, 3},
		{5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2, 0},
		{6, [][]int{{0, 3, 5}, {2, 3, 7}, {0, 5, 2}, {0, 2, 5}, {1, 2, 6}, {1, 4, 7}, {3, 4, 4}, {2, 5, 5}, {1, 5, 8}}, 8279, 5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findTheCity(ts.n, ts.edges, ts.distanceThreshold))
		})
	}
}
