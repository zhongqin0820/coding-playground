package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	p := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%f", &p[i])
	}
	fmt.Println(solve(N, p))
}

func solve(n int, p []float64) float64 {
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	dp[0][0] = 1
	//
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] * (1 - p[i-1])
				continue
			}
			dp[i][j] = dp[i-1][j]*(1-p[i-1]) + dp[i-1][j-1]*p[i-1]
		}
	}
	//
	var ret float64
	for j := n/2 + 1; j <= n; j++ {
		ret += dp[n][j]
	}
	return ret
}
