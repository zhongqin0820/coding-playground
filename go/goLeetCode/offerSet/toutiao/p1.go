package xmain

import (
	"fmt"
)

func JudgeN(n int) bool {
	if n <= 0 || n > 1024 {
		return false
	}
	return true
}

func DivMod(a, b int) (int, int) {
	return a / b, a % b
}

func xmain() {
	N := 0
	fmt.Scan(&N)
	res := 0
	if JudgeN(N) == false {
		fmt.Println(res)
		return
	}
	if N%1024 == 0 {
		fmt.Println(res)
		return
	}
	i, A, B := 0, 1024, 64
	i, N = DivMod(A-N, B)
	res += i
	A, B = 64, B>>2
	for N != 0 {
		i, N = DivMod(N, B)
		A = B
		B = B >> 2
		res += i
	}
	fmt.Println(res)
}
