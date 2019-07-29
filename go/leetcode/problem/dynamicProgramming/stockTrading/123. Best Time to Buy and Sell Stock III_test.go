package problem

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/

// Say you have an array for which the i^th element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete at most two transactions.

// Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

// max = max(prices[:i]) + max(prices[i:])
func maxProfitIII(prices []int) int {
	firstBuy, firstSell := math.MinInt32, 0
	secondBuy, secondSell := math.MinInt32, 0
	for _, curP := range prices {
		if firstBuy < -curP {
			firstBuy = -curP
		}
		if firstSell < firstBuy+curP {
			firstSell = firstBuy + curP
		}
		if secondBuy < firstSell-curP {
			secondBuy = firstSell - curP
		}
		if secondSell < secondBuy+curP {
			secondSell = secondBuy + curP
		}
	}
	return secondSell
}

func TestMaxProfitIII(t *testing.T) {
	tests := []struct {
		prices []int
		output int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
		// Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
		{[]int{1, 2, 3, 4, 5}, 4},
		// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
		// Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
		// engaging multiple transactions at the same time. You must sell before buying again.
		{[]int{7, 6, 4, 3, 1}, 0},
		// Explanation: In this case, no transaction is done, i.e. max profit = 0.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProfitIII(ts.prices))
		})
	}
}
