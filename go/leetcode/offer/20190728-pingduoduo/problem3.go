package main

import (
	"fmt"
	"log"
)

func problem3(n int, m int, P, A, B []int) []int {
	res := []int{}

	return res
}

func problem3Wrapper() {
	n, m, temp1, temp2 := 0, 0, 0, 0
	P, A, B := make([]int, 0, 1000), make([]int, 0, 50000), make([]int, 0, 50000)
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp1)
		P = append(P, temp1)
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &temp1, &temp2)
		A = append(A, temp1)
		B = append(B, temp2)
	}
	log.Println(n, m, P, A, B)
	fmt.Println(problem3(n, m, P, A, B))
}

// func main() {
// 	problem3Wrapper()
// }
