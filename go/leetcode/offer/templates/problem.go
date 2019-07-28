package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// replace the inputs & output
func problem(A, B []int) []int {
	res := []int{}
	return res
}

// replace the inputs & output
func problemWrapper() {
	// dealing with inputs
	A, B := input()
	// print problemWrpper()
	output := problem(A, B)
	fmt.Println(output)
}

// replace the inputs
func input() ([]int, []int) {
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

// the main entrance
func main() {
	problemWrapper()
}
