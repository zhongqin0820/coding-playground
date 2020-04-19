package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(solve(N, a))
}

func solve(n int, a []int) float64 {
	ones, twos, threes := 0, 0, 0
	for _, num := range a {
		switch num {
		case 1:
			ones++
		case 2:
			twos++
		case 3:
			threes++
		}
	}
	max_ones := ones + twos + threes
	max_twos := twos + threes
	max_threes := threes
	//
	dp := make([][][]float64, n+1)
	for i := range dp {
		dp[i] = make([][]float64, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n+1)
		}
	}
	//
	var i, j, k int
	var sum float64
	for k = 0; k <= max_threes; k++ {
		for j = 0; j <= max_twos-k; j++ {
			for i = 0; i <= max_ones-j-k; i++ {
				if i+j+k == 0 {
					continue
				}
				sum = float64(i + j + k)
				dp[i][j][k] = float64(n) / sum
				if i > 0 {
					dp[i][j][k] += (dp[i-1][j][k] * float64(i)) / sum
				}
				if j > 0 {
					dp[i][j][k] += (dp[i+1][j-1][k] * float64(j)) / sum
				}
				if k > 0 {
					dp[i][j][k] += (dp[i][j+1][k-1] * float64(k)) / sum
				}
			}
		}
	}
	return dp[i-1][j-1][k-1]
}
