package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ref: https://leetcode.com/problems/jump-game-iv/discuss/502699/Java-BFS-Solution-Clean-code-O(N)
func minJumps(arr []int) int {
	size := len(arr)
	m := make(map[int][]int, size)
	for i, a := range arr {
		m[a] = append(m[a], i)
	}
	dp := make([]int, size)
	for i := 0; i < size-1; i++ {
		dp[i] = -1
	}
	queue := []int{size - 1}
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		next := append(m[arr[i]], []int{i - 1, i + 1}...)
		for _, j := range next {
			if j >= 0 && j < size && dp[j] == -1 {
				dp[j] = dp[i] + 1
				queue = append(queue, j)
			}
		}
		delete(m, arr[i])
	}
	return dp[0]
}

func TestMinJumps(t *testing.T) {
	tests := []struct {
		arr    []int
		output int
	}{
		{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3},
		{[]int{7}, 0},
		{[]int{7, 6, 9, 6, 9, 6, 9, 7}, 1},
		{[]int{6, 1, 9}, 2},
		{[]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minJumps(ts.arr))
		})
	}
}
