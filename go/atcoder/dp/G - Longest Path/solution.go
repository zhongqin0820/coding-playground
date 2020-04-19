package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)
	graph := make(map[int][]int, N)
	var x, y int
	for i := 0; i <= M; i++ {
		fmt.Scanf("%d %d\n", &x, &y)
		if graph[x] == nil {
			graph[x] = make([]int, 0)
		}
		graph[x] = append(graph[x], y)
	}
	fmt.Println(solve(N, graph))
}

func solve(n int, g map[int][]int) int {
	res := 0
	dp := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		res = max(res, dfs(i, g, dp))
	}
	return res
}

// 使用dp[i] 记忆化dfs(i) 减少计算
func dfs(i int, g map[int][]int, dp map[int]int) int {
	if v, ok := dp[i]; ok {
		return v
	}
	temp := 0
	for _, nv := range g[i] {
		temp = max(temp, dfs(nv, g, dp)+1)
	}
	dp[i] = temp
	return temp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
