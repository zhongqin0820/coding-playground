package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jump-game-iii/
func canReach(arr []int, start int) bool {
	// 可达问题使用bfs
	size := len(arr)
	visited := make(map[int]bool, size)
	q := make([]int, 1, size)
	q[0] = start
	for len(q) != 0 {
		i := q[0]
		q = q[1:]
		if arr[i] == 0 {
			return true
		}
		if _, ok := visited[i]; ok {
			continue
		} else {
			visited[i] = true
		}
		if i+arr[i] < size {
			q = append(q, i+arr[i])
		}
		if i-arr[i] >= 0 {
			q = append(q, i-arr[i])
		}
	}

	return false
}

func TestCanReach(t *testing.T) {
	tests := []struct {
		arr    []int
		start  int
		output bool
	}{
		{[]int{4, 2, 3, 0, 3, 1, 2}, 5, true},
		{[]int{4, 2, 3, 0, 3, 1, 2}, 0, true},
		{[]int{3, 0, 2, 1, 2}, 2, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canReach(ts.arr, ts.start))
		})
	}
}
