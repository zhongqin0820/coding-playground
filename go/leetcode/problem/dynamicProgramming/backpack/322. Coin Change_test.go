package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/coin-change/description/

// You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

// Note:
// You may assume that you have an infinite number of each kind of coin.

func coinChange(coins []int, amount int) int {
	if amount == 0 || coins == nil || len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if i == coin {
				dp[i] = 1
			} else if dp[i] == 0 && dp[i-coin] != 0 {
				dp[i] = dp[i-coin] + 1
			} else if dp[i-coin] != 0 {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		output int
	}{
		{[]int{1, 2, 5}, 11, 3},
		// Explanation: 11 = 5 + 5 + 1
		{[]int{2}, 3, -1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, coinChange(ts.coins, ts.amount))
		})
	}
}
