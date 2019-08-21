package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem1(s string) string {
	A := strings.Split(s, " ")
	B := []string{""}
	n := len(A)
	for i := 1; i < n; i++ {
		if strings.Compare(A[i], "A") == 0 {
			B = append(B, []string{"12", "34"}...)
		} else if strings.Compare(A[i], "B") == 0 {
			B = append(B, []string{"AB", "CD"}...)
		} else {
			B = append(B, A[i])
		}
	}
	B[0] = strings.ToUpper(strconv.FormatInt(int64(len(B)), 16))
	return strings.Join(B, " ")
}

func problemWrapper1() {
	A := input1()
	output := problem1(A)
	fmt.Println(output)
}

func input1() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	return text
}

func main1() {
	problemWrapper1()
}
