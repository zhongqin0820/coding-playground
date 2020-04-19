package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	stones := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &stones[i])
	}
	fmt.Println(solve(n, k, stones))
}

func solve(n int, K int, arr []int) int {
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(arr[1] - arr[0])
	for i := 2; i < n; i++ {
		dp[i] = math.MaxInt32
		for k := 1; k <= min(i, K); k++ {
			dp[i] = min(dp[i], dp[i-k]+abs(arr[i]-arr[i-k]))
		}
	}
	return dp[n-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
