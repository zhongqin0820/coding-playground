package contestSet

import (
	// "container/list"
	"fmt"
	"log"
)

func prob873() {
	// N := 2
	// var dp [][]int = make([][]int, N*N)
	// var m map[int]int = make(map[int]int)
	// m[1] = 1
	// log.Printf("%v", m)
	// log.Printf("%q", dp)
	res := lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8})
	log.Println(res)
}

func lenLongestFibSubseq(A []int) int {
	if A == nil || len(A) < 3 || len(A) > 1000 {
		return -1
	}
	var m map[int]int = make(map[int]int)
	var N, res, k = len(A), -1, -1
	var dp [1000][1000]int
	for j := 0; j < N; j++ {
		m[A[j]] = j
		for i := 0; i < j; i++ {
			if v, ok := m[A[j]-A[i]]; ok {
				k = v
			} else {
				k = -1
			}
			// log.Println(k)
			if A[j]-A[i] < A[i] && k >= 0 {
				dp[i][j] = dp[k][i] + 1
			} else {
				dp[i][j] = 2
			}
			// log.Println(dp)
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	if res > 2 {
		return res
	}
	return 0
}

func main2() {
	prob873()
	fmt.Println("...")
}
