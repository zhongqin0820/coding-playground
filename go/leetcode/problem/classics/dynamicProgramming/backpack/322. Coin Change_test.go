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

// dp[0] = 0 // dp[i] 表示组成金额为i的最少硬币数
// dp[i] = amount + 1
// dp[i] = min(dp[i], dp[i-coins[j]] + 1)
func coinChangeII(coins []int, amount int) int {
	if amount == 0 || coins == nil || len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin >= 0 && dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
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
			ast.Equal(ts.output, coinChangeII(ts.coins, ts.amount))
		})
	}
}
