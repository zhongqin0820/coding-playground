package main

import (
	"fmt"
	"sort"
)

var mm map[byte]int = make(map[byte]int, 26)

func problem1(n int, s str) str {
	sort.Sort(s)
	return s
}

func problemWrapper1() {
	B, C := input1()
	output := problem1(B, C)
	for _, out := range output {
		fmt.Println(out)
	}
}

type str []string

func (s str) Len() int {
	return len(s)
}

func (s str) Less(i, j int) bool {
	li, lj := len(s[i]), len(s[j])
	if li < lj {
		for k, _ := range s[i] {
			// 如果二者下标不相等
			if s[i][k] != s[j][k] {
				return mm[s[i][k]] < mm[s[j][k]]
			}
		}
	} else {
		for k, _ := range s[j] {
			// 如果二者下标不相等
			if s[i][k] != s[j][k] {
				return mm[s[i][k]] < mm[s[j][k]]
			}
		}
	}
	return li < lj
}

func (s str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func input1() (int, str) {
	var m string
	fmt.Scanf("%s", &m)
	for i := range m {
		mm[m[i]] = i
	}
	var n int
	fmt.Scanf("%d", &n)
	var s str = make(str, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s[i])
	}
	return n, s
}

// func main() {
// 	problemWrapper1()
// }
