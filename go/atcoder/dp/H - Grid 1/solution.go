package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)
	scanner := bufio.NewScanner(os.Stdin)
	g := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		g = append(g, []byte(line))
	}
	fmt.Println(solve(N, M, &g))
}

func solve(n int, m int, g *[][]byte) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[1][1] = 1
	// dp[i][j] = dp[i-1][j]+dp[i][j-1]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == 1 && j == 1 {
				continue
			} else if (*g)[i-1][j-1] == '.' {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % (1e9 + 7)
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
