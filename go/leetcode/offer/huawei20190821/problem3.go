package main

import (
	"fmt"
	"strings"
)

func problem3(A string, B int, C [][]string) int {
	var M map[string][]string = make(map[string][]string)
	for _, c := range C {
		for i, s := range c {
			M[s] = append(M[s], c[:i]...)
			M[s] = append(M[s], c[i+1:]...)
		}
	}
	res := map[string]struct{}{}
	if _, ok := M[A]; !ok {
		return 0
	} else {
		queue := []string{A}
		for len(queue) > 0 {
			s := queue[0]
			queue = queue[1:]
			res[s] = struct{}{}
			if c, ok := M[s]; ok {
				for _, ss := range c {
					if _, ok := res[ss]; !ok {
						queue = append(queue, ss)
					}
				}
			}
		}
	}
	return len(res)
}

func problem3Wrapper() {
	A, B, C := input3()
	output := problem3(A, B, C)
	fmt.Print(output)
}

func input3() (string, int, [][]string) {
	var A string
	var B int
	fmt.Scanf("%s", &A)
	fmt.Scanf("%d", &B)
	var C [][]string
	for i := 0; i < B; i++ {
		var s string
		fmt.Scanf("%s", &s)
		c := strings.Split(s, ",")
		C = append(C, c)
	}
	return A, B, C
}

func main() {
	problem3Wrapper()
}
