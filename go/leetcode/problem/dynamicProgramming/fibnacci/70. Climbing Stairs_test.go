package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/climbing-stairs/description/

// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Note: Given n will be a positive integer.

func climbStairsTimeLimitExceeded(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	pre, cur := 0, 1
	for i := 0; i < n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{2, 2},
		// Explanation: There are two ways to climb to the top.
		// 1. 1 step + 1 step
		// 2. 2 steps
		{3, 3},
		// Explanation: There are three ways to climb to the top.
		// 1. 1 step + 1 step + 1 step
		// 2. 1 step + 2 steps
		// 3. 2 steps + 1 step
		{45, 1836311903},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, climbStairs(ts.n))
		})
	}
}
