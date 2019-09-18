package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-town-judge/
func findJudge(N int, trust [][]int) int {
	// trust[0]边的起点，trust[1]边的终点
	// count是入度统计数组，trustSomebody是是否相信别人的标志数组
	count := [1001]int{}
	trustSomebody := [1001]bool{}
	for _, t := range trust {
		trustSomebody[t[0]] = true
		count[t[1]]++
	}
	// 符合答案的要求：入度为N-1的节点且不相信别人
	for i := 1; i <= N; i++ {
		if count[i] == N-1 && !trustSomebody[i] {
			return i
		}
	}

	return -1
}

func TestFindJudge(t *testing.T) {
	tests := []struct {
		N      int
		trust  [][]int
		output int
	}{
		{2, [][]int{[]int{1, 2}}, 2},
		{3, [][]int{[]int{1, 3}, []int{2, 3}}, 3},
		{3, [][]int{[]int{1, 3}, []int{2, 3}, []int{3, 1}}, -1},
		{3, [][]int{[]int{1, 2}, []int{2, 3}}, -1},
		{4, [][]int{[]int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{4, 3}}, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findJudge(ts.N, ts.trust))
		})
	}
}
