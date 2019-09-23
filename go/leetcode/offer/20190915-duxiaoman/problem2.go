package main

import (
	"fmt"
)

func problem2(N, A, B, C int, a []int) int {
	res := 0
	return res
}

func problemWrapper2() {
	N, A, B, C, a := input2()
	output := problem2(N, A, B, C, a)
	fmt.Println(N, A, B, C, a)
	fmt.Println(output)
}

func input2() (N, A, B, C int, a []int) {
	fmt.Scanf("%d %d %d %d", &N, &A, &B, &C)
	a = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}
	return
}

// 7 1 1 1
// 3 6 4 3 4 5 6
// func main() {
// 	problemWrapper2()
// }
