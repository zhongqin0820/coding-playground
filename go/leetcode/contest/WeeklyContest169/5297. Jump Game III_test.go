package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-169/problems/jump-game-iii
func canReach(arr []int, start int) bool {
	return helper5297(arr, start, len(arr))
}

func helper5297(a []int, i int, n int) bool {
	if i >= n || i < 0 || (i == a[i] && a[0] != 0) {
		return false
	}
	if a[i] == 0 {
		return true
	}
	return helper5297(a, i-a[i], n) || helper5297(a, i+a[i], n)
}

func canReachII(arr []int, start int) bool {
	// 广度优先解决可达问题
	n := len(arr)
	visited := make(map[int]bool, n)
	q := make([]int, 0, n)
	q = append(q, start)
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
		if i+arr[i] < n {
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
			ast.Equal(ts.output, canReachII(ts.arr, ts.start))
		})
	}
}
