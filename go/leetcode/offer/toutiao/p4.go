package main

import (
	"fmt"
)

func MergeSort(s []int, flag bool) []int {
	Len := len(s)
	if Len <= 1 {
		return s
	}
	Mid := Len / 2
	a := MergeSort(s[:Mid], flag)
	b := MergeSort(s[Mid:], flag)
	return Merge(a, b, flag)
}

func Merge(a, b []int, flag bool) (ret []int) {
	var i, j int
	if !flag {
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] < b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
	} else {
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] > b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
	}
	ret = append(ret, a[i:]...)
	ret = append(ret, b[j:]...)
	return ret
}
func main() {
	N, M := 0, 0
	fmt.Scan(&N, &M)
	a := []int{}
	temp := 0
	for i := N; i > 0; i-- {
		fmt.Scan(&temp)
		a = append(a, temp)
	}
	b := MergeSort(a, false)
	// fmt.Println(b)
	var f = float32(b[len(b)-1])   //biggest
	var tempF = float32(M - N + 1) //rest
	// fmt.Println(f, tempF)
	fmt.Printf("%.2f \n", f/tempF)
}
