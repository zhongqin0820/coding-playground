package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func problem1(A, B []int) []int {
	if len(A) <= 1 {
		return A
	}
	sort.Ints(B)
	idx := 0
	for i := 1; i < len(A); i++ {
		if A[idx] < A[i] {
			idx = i
			continue
		}
		break
	}
	idx++
	n := len(B)
	i := sort.Search(n, func(i int) bool {
		return A[idx-1] < B[i] && B[i] < A[idx+1]
	})
	for j := i; j < n; j++ {
		if A[idx-1] < B[j] && B[j] < A[idx+1] {
			i = j
		} else {
			break
		}
	}
	if i == n {
		return []int{}
	}
	A[idx] = B[i]
	return A
}

func problem1Wrapper() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	texts := strings.Split(text, " ")
	A := []int{}
	for _, v := range texts {
		num, _ := strconv.Atoi(v)
		A = append(A, num)
	}
	B := []int{}
	text, _ = reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	texts = strings.Split(text, " ")
	for _, v := range texts {
		num, _ := strconv.Atoi(v)
		B = append(B, num)
	}
	C := problem1(A, B)
	if len(C) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(strings.Trim(fmt.Sprintf("%v", A), "[]"))
	}
}

// func main() {
// 	problem1Wrapper()
// }
