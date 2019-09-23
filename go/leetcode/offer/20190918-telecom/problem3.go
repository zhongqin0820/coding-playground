package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem3(A []int) int {
	res, cur := 0, 0
	for _, v := range A {
		if v+cur > 0 {
			cur += v
			if cur > res {
				res = cur
			}
		} else {
			cur = 0
		}
	}
	return res
}

func problemWrapper3() {
	A := input3()
	output := problem3(A)
	fmt.Println(output)
}

func input3() []int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "]\n")
	text = strings.TrimLeft(text, "[")
	texts := strings.Split(text, ", ")
	A := []int{}
	for _, v := range texts {
		num, _ := strconv.Atoi(v)
		A = append(A, num)
	}
	return A
}

// [2, 4, -2, 5, -6]
func main() {
	problemWrapper3()
}
