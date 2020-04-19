package main

import (
	"fmt"
)

const mod = 1e9 + 7

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(solve(n, k, a))
}

func solve(N, K int, a []int) int {
	// dp[n][k] stands for the ways of giving k candies to n children
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	dp[0][0] = 1
	//
	for i := 1; i <= N; i++ {
		dp[i][0] = 1
		for j := 0; j < K; j++ {
			dp[i][j+1] = (dp[i][j] + dp[i-1][j+1]) % mod
			if j-a[i-1] >= 0 { // 当前的糖数满足条件
				dp[i][j+1] -= dp[i-1][j-a[i-1]]
			}
		}
	}
	return (dp[N][K]) % mod
}
