package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func problem1(A [][]int, n, m int) int {
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				continue
			}
			flag := true
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				if 0 <= x && x < n && 0 <= y && y < m && A[x][y] == 1 {
					continue
				}
				flag = false
				break
			}
			if flag {
				res++
			}
		}
	}

	return res
}

func problemWrapper1() {
	A, n, m := input1()
	output := problem1(A, n, m)
	fmt.Println(output)
}

func input1() ([][]int, int, int) {
	reader := bufio.NewReader(os.Stdin)
	n, m := 0, 0
	fmt.Scanf("%d %d", &n, &m)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		texts := strings.Split(text, "")
		for _, v := range texts {
			num, _ := strconv.Atoi(v)
			A[i] = append(A[i], num)
		}
	}
	return A, n, m
}

// func main() {
// 	problemWrapper1()
// }
