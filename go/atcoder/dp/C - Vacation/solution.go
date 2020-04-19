package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	activities := make([][]int, 0)
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := scanner.Text()
		temps := strings.Split(temp, " ")
		day := make([]int, 4)
		m := 1
		for i, str := range temps {
			num, _ := strconv.Atoi(str)
			day[i] = num
			m = max(m, num)
		}
		day[3] = m
		activities = append(activities, day)
	}
	fmt.Println(solve(N, &activities))
}

// dp[i][j] 表示第i天选择活动j得到的 the maximum possible total points of happiness
// dp[i][j] = max_{{j=0}^2, {k=0}^2}(dp[i][j], dp[i-1][k!=j])
// dp[i][j] = max_{j=0}^2(dp[i][j], max_{{k=0}^2}(dp[i-2][k]))

func solve(n int, arr *[][]int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	for k := 0; k < 4; k++ {
		dp[0][k] = (*arr)[0][k]
	}
	if n == 1 {
		return dp[0][3]
	}
	//
	m := 1
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			if j == k {
				continue
			}
			dp[1][j] = max(dp[1][j], dp[0][k]+(*arr)[1][j])
		}
		m = max(m, (*arr)[1][j])
	}
	dp[1][3] = m
	//
	for i := 2; i < n; i++ {
		m = 1
		for j := 0; j < 3; j++ {
			m = max(m, (*arr)[i][j])
			dp[i][j] = (*arr)[i][j] + dp[i-2][3] // dp[i] = arr[i] + max(arr[i-2])
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				// dp[i] = arr[i] + max_{k!=j}(dp[i-1][k])
				dp[i][j] = max(dp[i][j], dp[i-1][k]+(*arr)[i][j])
			}
		}
		dp[i][3] = m
	}
	// 返回结果
	m = 1
	for _, cnt := range dp[n-1] {
		m = max(m, cnt)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
