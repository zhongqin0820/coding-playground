package main

import (
	"fmt"
)

// Q4
func judge(a, b int) bool {
	if a < 1 || a > 1000000 {
		return false
	}
	if b < 0 || b > 2000 {
		return false
	}
	return true
}

func judgeNum(n int) bool {
	if n < 1 || n > 5 {
		return false
	}
	return true
}

func main() {
	n, s := 0, 0
	fmt.Scan(&n, &s)
	if judge(n, s) == false {
		return
	}
	a := []int{}
	var temp int
	for ; n > 0; n-- {
		fmt.Scan(&temp)
		if judgeNum(temp) == false {
			return
		}
		a = append(a, temp)
	}
	f := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	var i, j = 0, 0
	res := 0
	for ; i < len(a); i++ {
		if a[i] == 0 {
			j = i + 1
		} else if a[i] == a[j] {
			j++
		}
		if c, _ := f[a[i]]; c == 0 {
			f[a[i]] = 1
		} else {
			res++
		}
		if res == 5 {
			fmt.Println(i - j)
		}
	}
	fmt.Println(-1)
}
