package problem

import (
	"fmt"
)

func Print1ToN(n int) {
	arr := make([]byte, n)
	str := []byte("0123456789")
	for i := 0; i < 10; i++ {
		arr[0] = str[i]
		perm(arr, n, 1)
	}
}

func perm(arr []byte, n, i int) {
	if i >= n {
		printNum(arr, n)
		return
	}
	str := []byte("0123456789")
	for j := 0; j < 10; j++ {
		arr[i] = str[j]
		perm(arr, n, i+1)
	}
}

func printNum(arr []byte, n int) {
	sFlag := true
	for i := 0; i < n; i++ {
		if sFlag && arr[i] == '0' {
			continue
		}
		sFlag = false
		fmt.Printf("%c", arr[i])
	}
	fmt.Print(" ")
}
