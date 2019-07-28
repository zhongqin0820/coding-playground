package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func problem2(s string) bool {
	m := make(map[byte]int, 26)
	ss := strings.Split(s, " ")
	for i := 0; i < len(ss); i++ {
		j, k := 0, len(ss[i])-1
		m[ss[i][j]]++
		m[ss[i][k]]++
	}
	flag := 0
	for _, v := range m {
		if v%2 != 0 {
			flag++
			if flag == 2 {
				break
			}
		}
	}
	return flag < 2
}

func problem2Wrapper() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n")
	fmt.Println(problem2(text))
}

// func main() {
// 	problem2Wrapper()
// }
