package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/

// Say you have an array for which the ith element is the price of a given stock on day i.

// Design an algorithm to find the maximum profit. You may complete at most k transactions.

// Note:
// You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

func maxProfitIV(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	if k >= size {
		return helper188(prices)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	local := make([]int, k+1)
	global := make([]int, k+1)

	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(diff, 0), local[j]+diff)
			global[j] = max(local[j], global[j])
		}
	}

	return global[k]
}

func helper188(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}

func TestMaxProfitIV(t *testing.T) {
	tests := []struct {
		k      int
		prices []int
		output int
	}{
		{2, []int{2, 4, 1}, 2},
		// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
		// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
		// Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxProfitIV(ts.k, ts.prices))
		})
	}
}
