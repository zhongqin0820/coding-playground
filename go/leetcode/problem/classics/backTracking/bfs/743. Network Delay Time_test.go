package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func networkDelayTime(times [][]int, N int, K int) int {
	maxInt := int(1e4)

	// minTime[i] == m 表示， i 节点接收到信号所需的最小时间为 m
	minTime := make([]int, N+1)
	for i := 1; i <= N; i++ {
		minTime[i] = maxInt
	}
	// 信号从 K 节点出发，所以 minTime[K] = 0
	minTime[K] = 0

	// cost[i][j] = m 表示，从 i 节点到 j 节点所需的时间
	cost := make([][]int, N+1)
	for i := range cost {
		cost[i] = make([]int, N+1)
		for j := 0; j <= N; j++ {
			cost[i][j] = maxInt
		}
	}
	for _, t := range times {
		cost[t[0]][t[1]] = t[2]
	}

	queue := make([]int, 1, N+1)
	queue[0] = K
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 1; v <= N; v++ {
			if minTime[v] > minTime[u]+cost[u][v] {
				minTime[v] = minTime[u] + cost[u][v]
				queue = append(queue, v)
			}
		}
	}

	res := 0
	for i := 1; i <= N; i++ {
		res = helper743max(res, minTime[i])
	}

	if res == maxInt {
		// 存在无法到达的节点
		return -1
	}
	return res
}

func helper743max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		times  [][]int
		N      int
		K      int
		output int
	}{
		{
			[][]int{{2, 1, 1}},
			2, 2, 1,
		},
		{
			[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4, 2, 2,
		},
		{
			[][]int{{2, 1, 3}, {2, 3, 1}, {3, 4, 1}},
			4, 2, 3,
		},
		{
			[][]int{
				{3, 5, 78},
				{2, 1, 1},
				{1, 3, 0},
				{4, 3, 59},
				{5, 3, 85},
				{5, 2, 22},
				{2, 4, 23},
				{1, 4, 43},
				{4, 5, 75},
				{5, 1, 15},
				{1, 5, 91},
				{4, 1, 16},
				{3, 2, 98},
				{3, 4, 22},
				{5, 4, 31},
				{1, 2, 0},
				{2, 5, 4},
				{4, 2, 51},
				{3, 1, 36},
				{2, 3, 59}},
			5, 5, 31,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, networkDelayTime(ts.times, ts.N, ts.K))
		})
	}
}
