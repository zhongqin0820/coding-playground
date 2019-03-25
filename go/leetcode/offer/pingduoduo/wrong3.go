package main

import (
	"fmt"
)

func judgeN(n int) bool {
	if n < 1 || n > 20000 {
		return false
	}
	return true
}

func judgeD(n int) bool {
	if n < 1 || n > 10000000 {
		return false
	}
	return true
}

func judgeX(n int) bool {
	if n < 1 || n > 10000000 {
		return false
	}
	return true
}

func main() {
	n, d := 0, 0
	var a, b []int
	fmt.Scan(&n, &d)
	if judgeN(n) == false || judgeD(d) == false {
		return
	}
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		if judgeX(x) == false || judgeX(y) == false {
			return
		}
		a = append(a, x)
		b = append(b, y)
	}
	var max = 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[j]-a[i] > d && b[j]+b[i] > max {
				max = b[j] + b[i]
			}
		}
	}
	fmt.Println(max)
}
