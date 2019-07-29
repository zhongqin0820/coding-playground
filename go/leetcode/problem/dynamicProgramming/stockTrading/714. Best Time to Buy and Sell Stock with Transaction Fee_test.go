package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
// You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
// Return the maximum profit you can make.

// Note:
// 1. 0 < prices.length <= 50000.
// 1. 0 < prices[i] < 50000.
// 1. 0 <= fee < 50000.

func maxProfitII(prices []int, fee int) int {
	empty := 0
	hold := -1 << 63
	for _, p := range prices {
		temp := empty
		empty = helper714(empty, hold+p)
		hold = helper714(hold, temp-p-fee)
	}
	return empty
}

func helper714(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		prices []int
		fee    int
		output int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		// Explanation: The maximum profit can be achieved by:
		// Buying at prices[0] = 1
		// Selling at prices[3] = 8
		// Buying at prices[4] = 4
		// Selling at prices[5] = 9
		// The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProfitII(ts.prices, ts.fee))
		})
	}
}
