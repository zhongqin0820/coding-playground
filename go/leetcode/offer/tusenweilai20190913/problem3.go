package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func problem3(A, B []int) []int {
	res := []int{}
	return res
}

func problemWrapper3() {
	A, B := input3()
	output := problem3(A, B)
	fmt.Println(output)
}

func input3() ([]int, []int) {
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

	log.Println(A, B) // for check
	return A, B
}

// func main() {
// 	problemWrapper3()
// }
