package main

import (
	"fmt"
	"math"
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

// dp[i][v] = min(dp[i-1][v], dp[i-1][v-vs[i]]+w[i])
func solve(n, W int, w, vs []int) int {
	dp := make([][]int, n+1)
	// 注意初始化条件
	V := n * 1000 // 1 ≤ v_i ≤ 1000
	for i := range dp {
		dp[i] = make([]int, V+1)
	}
	for v := 1; v <= V; v++ {
		dp[0][v] = math.MaxInt32
	}
	// 特别注意不足的情况的处理
	for i := 1; i <= n; i++ {
		for v := 0; v <= V; v++ {
			if v < vs[i-1] {
				dp[i][v] = dp[i-1][v]
			} else {
				dp[i][v] = min(dp[i-1][v], dp[i-1][v-vs[i-1]]+w[i-1])
			}
		}
	}
	// 注意返回结果的条件
	res := 0
	for v := 0; v <= V; v++ {
		if dp[n][v] <= W {
			res = v
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
