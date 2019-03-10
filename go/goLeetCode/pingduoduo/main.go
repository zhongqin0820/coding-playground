package main

import (
	"container/list"
	"fmt"
	// "strings"
)

func judgeN(n string) bool {
	if len(n) == 0 || len(n) > 2500 {
		return false
	}
	var flag = true
	for _, v := range []byte(n) {
		if v == '(' || v == ')' {
			continue
		}
		flag = false
		break
	}
	return flag
}

func judgeB(b []byte) bool {
	l := list.New()
	if b[0] == ')' {
		return false
	}
	l.PushFront(b[0])
	// fmt.Println(l.Len())
	for i := 1; i < len(b); i++ {
		if b[i] == ')' && l.Len() > 0 {
			l.Remove(l.Front())
			// fmt.Println(l.Len())
		} else {
			l.PushFront(b[i])
			// fmt.Println(l.Len())
		}
	}
	// fmt.Println(l.Len())
	if l.Len() == 0 {
		return true
	}
	return false
}

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	if judgeN(a) == false || judgeN(b) == false || len(a) != len(b) {
		return
	}
	// fmt.Println("legal char")
	var c []byte
	c = append(c, []byte(a)...)
	c = append(c, []byte(b)...)
	var res int
	// fmt.Println(len(c))
	for i := 0; i < len(c)-1; i++ {
		for j := i + 1; j < len(c); j++ {
			c[i], c[j] = c[j], c[i]
			if judgeB(c) == true {
				res = res + 1
			}
		}
	}
	fmt.Println(res)
}
