package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(solve(n, k, a))
}

func solve(n, k int, a []int) string {
	dp := make([]bool, k+1)
	dp[0] = false
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			// 如果存在使对方必败的决策，则当前必胜
			if i-a[j] >= 0 && !dp[i-a[j]] {
				dp[i] = true
			}
		}
	}
	if dp[k] {
		return "First"
	}
	return "Second"
}
