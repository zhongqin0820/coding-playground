package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

// Say you have an array for which the i^th element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:
// You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
// After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

func maxProfit(prices []int) int {
	switch {
	case prices == nil || len(prices) == 0:
		return 0
	}
	n := len(prices)
	buy := make([]int, n+1)
	sel := make([]int, n+1)
	buy[1] = 0 - prices[0]
	for i := 2; i <= n; i++ {
		buy[i] = helper309(buy[i-1], sel[i-2]-prices[i-1])
		sel[i] = helper309(sel[i-1], buy[i-1]+prices[i-1])
	}
	return sel[n]
}

func helper309(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		output int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		// Explanation: transactions = [buy, sell, cooldown, buy, sell]
		{[]int{1}, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProfit(ts.prices))
		})
	}
}
