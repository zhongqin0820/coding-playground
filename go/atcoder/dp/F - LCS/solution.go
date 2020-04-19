package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scanf("%s\n%s\n", &s, &t)
	fmt.Println(string(solve([]byte(s), []byte(t))))
}

func solve(s, t []byte) []byte {
	// dp[i][j] 表示 s[:i]与t[:j]的maxLen of LCS
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	//
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// 打印恢复路径
	bs := make([]byte, dp[len(s)][len(t)])
	for i, j := len(s), len(t); dp[i][j] > 0; {
		if s[i-1] == t[j-1] {
			bs[dp[i][j]-1] = s[i-1]
			i--
			j--
		} else {
			if dp[i][j] == dp[i-1][j] {
				i--
			} else {
				j--
			}
		}
	}
	return bs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// xxxx i
// yyyyy j
