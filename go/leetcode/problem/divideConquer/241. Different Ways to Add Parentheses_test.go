package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/different-ways-to-add-parentheses/description/
func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)
	var dfs func(string) []int
	dfs = func(s string) []int {
		res := make([]int, 0, len(s))
		if t, ok := cache[s]; ok {
			return t
		}

		for i := range s {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				for _, left := range dfs(s[:i]) {
					for _, right := range dfs(s[i+1:]) {
						res = append(res, helper241(left, right, s[i]))
					}
				}
			}
		}

		if len(res) == 0 {
			temp, _ := strconv.Atoi(s)
			res = append(res, temp)
		}

		cache[s] = res
		return res
	}

	return dfs(input)
}

func helper241(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}

func TestDiffWaysToCompute(t *testing.T) {
	tests := []struct {
		input  string
		output []int
	}{
		{"2-1-1", []int{2, 0}},
		{"2*3-4*5", []int{-34, -10, -14, -10, 10}},
		{"2+2+2+2", []int{8, 8, 8, 8, 8}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, diffWaysToCompute(ts.input))
		})
	}
}
