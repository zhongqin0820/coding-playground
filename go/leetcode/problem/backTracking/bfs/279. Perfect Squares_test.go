package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/perfect-squares/description/

// Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}
	queue := []int{}
	marked := make([]bool, n+1)
	queue = append(queue, n)
	marked[n] = true

	for level := 1; len(queue) > 0; level++ {
		for size := len(queue); size > 0; size-- { //bfs
			cur := queue[0]
			queue = queue[1:]
			for _, s := range perfects { // 计算是否有符合的平方数
				next := cur - s
				if next < 0 {
					break
				}
				if next == 0 {
					return level
				}
				if marked[next] {
					continue
				}
				marked[next] = true
				queue = append(queue, next)
			}
		}
	}
	return n
}

func TestNumSquares(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{12, 3},
		// Explanation: 12 = 4 + 4 + 4.
		{13, 2},
		// Explanation: 13 = 4 + 9.
		{1, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numSquares(ts.n))
		})
	}
}
