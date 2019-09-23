package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	k := int(math.Sqrt(float64(n))) + 1
	for i := 5; i < k; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// replace the inputs & output
func problem(A, B int) int {
	res1, res2 := 0, 0
	for i := A; i < B; i++ {
		if isPrime(i) {
			if i/10 == 0 {
				res2 += i
			} else {
				res2 += i % 10
				res1 += (i % 100) / 10
			}
		}
	}
	if res1 < res2 {
		return res1
	}
	return res2
}

func problemWrapper() {
	A, B := input()
	output := problem(A, B)
	fmt.Println(output)
}

func input() (int, int) {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	return A, B
}

func main2() {
	problemWrapper()
}
