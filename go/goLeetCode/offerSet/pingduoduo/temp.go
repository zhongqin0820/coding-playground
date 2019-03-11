package main

import (
	"fmt"
)

func judgeN(n int) bool {
	if n > 50 || n < 0 {
		return false
	}
	return true
}

func judgeX(n int) bool {
	if n < 0 || n > 100 {
		return false
	}
	return true
}

// the order of s[k] and v decide the sorted order of the array
func QuickSort(s []int, flag bool) []int {
	if len(s) <= 1 {
		return []int{}
	}
	var i, j, v, k = 0, len(s) - 1, s[0], 1

	for i < j {
		if flag { //big->small
			if s[k] < v {
				s[k], s[j] = s[j], s[k]
				j--
			} else {
				s[i], s[k] = s[k], s[i]
				i++
				k++
			}
		} else { //small -> big
			if s[k] > v {
				s[k], s[j] = s[j], s[k]
				j--
			} else {
				s[i], s[k] = s[k], s[i]
				i++
				k++
			}
		}
	}
	QuickSort(s[:i], flag)
	QuickSort(s[i+1:], flag)
	return s
}

func main() {
	n := 0
	var a, b []int
	fmt.Scan(&n)
	if judgeN(n) == false {
		return
	}
	var x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if judgeX(x) == false {
			return
		}
		a = append(a, x)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if judgeX(x) == false {
			return
		}
		b = append(b, x)
	}
	a = QuickSort(a, true)
	b = QuickSort(b, false)
	res := 0
	for i := 0; i < n; i++ {
		res = res + a[i]*b[i]
	}
	fmt.Println(res)
}
