package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfitI(prices []int) int {
	max, temp := 0, 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}
	return max
}

func TestMaxProfitI(t *testing.T) {
	tests := []struct {
		prices []int
		output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProfitI(ts.prices))
		})
	}
}
