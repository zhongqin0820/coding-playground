package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	stones := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &stones[i])
	}
	fmt.Println(solve(N, stones))
}

func solve(N int, stones []int) int {
	dp := make([]int, N)
	dp[0] = 0
	dp[1] = abs(stones[1] - stones[0])
	for i := 2; i < N; i++ {
		// dp[i]: min_{j=1}^2(dp[i-j]+abs(arr[i]-arr[i-j]))
		dp[i] = min(dp[i-1]+abs(stones[i]-stones[i-1]), dp[i-2]+abs(stones[i]-stones[i-2]))
	}
	return dp[N-1]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
