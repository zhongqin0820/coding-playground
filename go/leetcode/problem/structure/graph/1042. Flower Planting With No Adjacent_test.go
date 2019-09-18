package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/flower-planting-with-no-adjacent/
// https://medium.com/@poitevinpm/solution-to-leetcode-problem-1042-flower-planting-with-no-adjacent-577e11cec28f
func gardenNoAdj(N int, paths [][]int) []int {
	// 将边构造图结构
	connects := make([][]int, N)
	for _, p := range paths {
		i, j := p[0]-1, p[1]-1 // i, j = x-1, y-1
		connects[i] = append(connects[i], j)
		connects[j] = append(connects[j], i)
	}
	res := make([]int, N)
	for i := 0; i < N; i++ {
		visited := [5]bool{}
		// 访问当前节点的所有邻边
		for _, j := range connects[i] {
			visited[res[j]] = true
		}
		// 上色
		for color := 1; color <= 4; color++ {
			if !visited[color] {
				res[i] = color
				break
			}
		}
	}
	return res
}

func TestGardenNoAdj(t *testing.T) {
	tests := []struct {
		N      int
		paths  [][]int
		output []int
	}{
		{
			3,
			[][]int{{1, 2}, {2, 3}, {3, 1}},
			[]int{1, 2, 3},
		},
		{
			4,
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 1, 2},
		},
		{
			4,
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}},
			[]int{1, 2, 3, 4},
		},
		{
			5,
			[][]int{{4, 1}, {4, 2}, {4, 3}, {2, 5}, {1, 2}, {1, 5}},
			[]int{1, 2, 1, 3, 3},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, gardenNoAdj(ts.N, ts.paths))
		})
	}
}
