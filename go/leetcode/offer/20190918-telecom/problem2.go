package main

import (
	"fmt"
	"strings"
)

func problem2(s string) bool {
	ss := strings.Split(s, ";")
	// 超过两个字符串
	if len(ss) != 2 {
		return false
	}
	// 长度不同
	if len(ss[0]) != len(ss[1]) {
		return false
	}
	n := len(ss[0])
	a, b := make(map[byte]int, n), make(map[byte]int, n)
	for i := range ss[0] {
		a[ss[0][i]]++
		b[ss[1][i]]++
		if a[ss[0][i]] != b[ss[1][i]] {
			return false
		}
	}
	return true
}

func problemWrapper2() {
	a := input2()
	output := problem2(a)
	if output {
		fmt.Println("True")
	} else {

		fmt.Println("False")
	}
}

func input2() (a string) {
	fmt.Scanf("%s", &a)
	return a
}

// func main() {
// 	problemWrapper2()
// }
