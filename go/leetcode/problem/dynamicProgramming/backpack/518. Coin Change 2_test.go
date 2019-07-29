package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/coin-change-2/description/

// You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.

// Note:
// You can assume that
// 0 <= amount <= 5000
// 1 <= coin <= 5000
// the number of coins is less than 500
// the answer is guaranteed to fit into signed 32-bit integer

func change(amount int, coins []int) int {
	if coins == nil || len(coins) == 0 {
		if amount == 0 {
			return 1
		}
		return 0
	}
	dp := make([]int, amount+1) //dp 记录可达成目标amount的组合数目。
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changeAdv(amount int, coins []int) int {
	// dp[i]表示总额为i时的方案数.
	// 转移方程: dp[i] = Σdp[i - coins[j]]; 表示 总额为i时的方案数 = 总额为i-coins[j]的方案数的加和.
	dp := make([]int, amount+1)
	// 记得初始化dp[0] = 1; 表示总额为0时方案数为1.
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ { // 从coin开始遍历，小于coin的值没有意义
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func TestChange(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		output int
	}{
		{5, []int{1, 2, 5}, 4},
		// Explanation: there are four ways to make up the amount:
		// 5=5
		// 5=2+2+1
		// 5=2+1+1+1
		// 5=1+1+1+1+1
		{3, []int{2}, 0},
		// Explanation: the amount of 3 cannot be made up just with coins of 2.
		{10, []int{10}, 1},
		{0, []int{}, 1},
		{7, []int{}, 0},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, change(ts.amount, ts.coins))
			ast.Equal(ts.output, changeAdv(ts.amount, ts.coins))
		})
	}
}
