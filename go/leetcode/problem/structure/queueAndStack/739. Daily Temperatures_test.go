package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/daily-temperatures/description/
func dailyTemperatures(T []int) []int {
	n := len(T)
	dp, flag := make([]int, n), false
	for i := n - 2; i >= 0; i-- {
		if T[i] < T[i+1] {
			dp[i] = 1
		} else {
			flag = false
			for j := i; j < n; j++ {
				dp[i] = j - i
				if T[i] < T[j] {
					flag = true
					break
				}
			}
			if !flag {
				dp[i] = 0
			}
		}
	}
	return dp
}

func dailyTemperaturesII(T []int) []int {
	n := len(T)
	res, stack, top := make([]int, n), make([]int, n), -1
	for i := 0; i < n; i++ {
		for top >= 0 && T[stack[top]] < T[i] {
			res[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}
	return res
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		T      []int
		output []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, dailyTemperatures(ts.T))
			ast.Equal(ts.output, dailyTemperaturesII(ts.T))
		})
	}
}
