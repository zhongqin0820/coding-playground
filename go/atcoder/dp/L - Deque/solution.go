package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(solve(n, a))
}

func solve(n int, a []int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// xxxxxi[xxxxxx]jxxxxxxx
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = a[i]
			} else {
				// dp[i][j] 表示a[i:j+1] 的先手最大收益
				dp[i][j] = max(a[i]-dp[i+1][j], a[j]-dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
