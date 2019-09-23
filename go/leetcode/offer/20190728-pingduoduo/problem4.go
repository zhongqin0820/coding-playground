package main

import (
	"fmt"
)

func problem4(n int, L, W []int) int {
	res := 0

	return res
}

func problem4Wrapper() {
	n, temp := 0, 0
	L, W := make([]int, 0, 1000), make([]int, 0, 1000)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		L = append(L, temp)
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		W = append(W, temp)
	}
	fmt.Println(problem4(n, L, W))
}

func main() {
	problem4Wrapper()
}
