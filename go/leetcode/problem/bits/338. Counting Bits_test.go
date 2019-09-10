package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/counting-bits/description/
func countBits(num int) []int {
	// dp[i] = dp[i&(i-1)] + 1;
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func TestCountBits(t *testing.T) {
	tests := []struct {
		num    int
		output []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countBits(ts.num))
		})
	}
}
