package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 第一行输入为长度N
// 接下来N行，每行一个数，代表数组A的N个元素
// 或者，接下来一行有N个数，用空格分割
func NCasesIntput1() (int, []int) {
	n := 0
	fmt.Scan(&n)

	MAXLEN := 1024 //如提供长度范围，替换1024
	res := make([]int, 0, MAXLEN)
	for i := 0; i < n; i++ {
		temp := 0
		fmt.Scan("%d", temp)
		res = append(res, temp)
	}
	return n, res
}

// 第一行输入为长度N
// 接下来N行，每行两个空格分割的数，分别对应A_i与B_i
func NCasesIntput2() (int, []int, []int) {
	n := 0
	fmt.Scan(&n)

	MAXLENA := 1024 //如提供长度范围，替换1024
	MAXLENB := 1024 //如提供长度范围，替换1024
	A := make([]int, 0, MAXLENA)
	B := make([]int, 0, MAXLENB)
	for i := 0; i < n; i++ {
		temp1, temp2 := 0, 0
		fmt.Scanf("%d %d", temp1, temp2)
		A = append(A, temp1)
		B = append(B, temp2)
	}
	return n, A, B
}

// 一行以空格间隔的输入，未知元素个数
func Input1(s string) []int {
	var input io.Reader
	if len(s) == 0 {
		input = os.Stdin
	} else {
		input = strings.NewReader(s)
	}
	reader := bufio.NewReader(input)

	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	texts := strings.Split(text, " ")
	MAXLEN := 1024 //如提供长度范围，替换1024
	A := make([]int, 0, MAXLEN)
	for _, v := range texts {
		num, _ := strconv.Atoi(v)
		A = append(A, num)
	}
	return A
}

// 一行以两个元素输入，未知行数
func Input2(s string) ([]int, []int) {
	var input io.Reader
	if len(s) == 0 {
		input = os.Stdin
	} else {
		input = strings.NewReader(s)
	}
	reader := bufio.NewReader(input) //input=os.Stdin
	//
	scanner := bufio.NewScanner(reader)
	//
	MAXLENA := 1024 //如提供长度范围，替换1024
	MAXLENB := 1024 //如提供长度范围，替换1024
	A := make([]int, 0, MAXLENA)
	B := make([]int, 0, MAXLENB)
	temp1, temp2 := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 || s == "\n" {
			break
		}
		if len(strings.Split(strings.TrimRight(s, " "), " ")) != 2 {
			break
		}
		fmt.Sscanf(s, "%d %d", &temp1, &temp2)
		A = append(A, temp1)
		B = append(B, temp2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return A, B
}

// 未知行数，未知列数，列间以空格
func Input3(s string) [][]int {
	var input io.Reader
	if len(s) == 0 {
		input = os.Stdin
	} else {
		input = strings.NewReader(s)
	}
	reader := bufio.NewReader(input)
	scanner := bufio.NewScanner(reader)
	MAXLEN := 1024
	A := make([][]int, 0, MAXLEN)
	var i int
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 || s == "\n" {
			break
		}
		tempS := strings.Split(strings.TrimRight(s, " "), " ")
		if len(tempS) == 0 {
			break
		}
		temp := []int{}
		for _, v := range tempS {
			num, _ := strconv.Atoi(v)
			temp = append(temp, num)

		}
		i++
		A = append(A, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	return A
}
