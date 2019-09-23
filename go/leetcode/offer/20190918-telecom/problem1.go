package main

import (
	"fmt"
)

func problem1(n int, A []int) int {
	res := A[0]
	for i := 1; i < n; i++ {
		res ^= A[i]
	}
	return res
}

func problemWrapper1() {
	n, A := input1()
	output := problem1(n, A)
	fmt.Println(output)
}

func input1() (n int, A []int) {
	fmt.Scanf("%d", &n)
	A = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &A[i])
	}
	return
}

// func main() {
// 	problemWrapper1()
// }
