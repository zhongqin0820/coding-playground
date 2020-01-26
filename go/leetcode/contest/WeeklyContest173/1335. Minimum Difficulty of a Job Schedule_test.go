package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minDifficulty(A []int, d int) int {
	if len(A) < d {
		return -1
	}
	// TODO[zoking](2020-01-26): ref: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/discuss/490323/Python-Top-down-and-Bottom-up-solutions-with-explanation
	return d
}

func TestMinDifficulty(t *testing.T) {
	// 1 <= jobDifficulty.length <= 300
	// 0 <= jobDifficulty[i] <= 1000
	// 1 <= d <= 10
	tests := []struct {
		A      []int
		d      int
		output int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 1, 7},
		{[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
		{[]int{7, 1, 7, 1, 7, 1}, 3, 15},
		{[]int{11, 111, 22, 222, 33, 333, 44, 444}, 6, 843},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minDifficulty(ts.A, ts.d))
		})
	}
}
