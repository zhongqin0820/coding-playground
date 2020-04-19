package main

import (
	"fmt"
)

func main() {
	var N, W int
	fmt.Scanf("%d %d\n", &N, &W)
	w, v := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d\n", &w[i], &v[i])
	}
	fmt.Println(solve(N, W, w, v))
}

// dp[i][c] = max(dp[i-1][c], dp[i-1][c-w[i]]+v[i])
// 特别注意容量不足的情况的处理
func solve(n, C int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, C+1)
	}
	for i := 1; i <= n; i++ {
		for c := 0; c <= C; c++ {
			if c-w[i-1] >= 0 {
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-w[i-1]]+v[i-1])
			} else {
				dp[i][c] = dp[i-1][c]
			}
		}
	}
	return dp[n][C]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
