package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func problem2(n int, nodes []int, edges map[int][]int) int {
	res := 0
	return res
}

func problemWrapper2() {
	n, nodes, edges := input2()
	output := problem2(n, nodes, edges)
	fmt.Println(n, nodes, edges)
	fmt.Println(output)
}

func input2() (int, []int, map[int][]int) {
	var n int
	fmt.Scanf("%d", &n)
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nodes[i])
	}
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	var edges map[int][]int = make(map[int][]int, n)
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
		if _, ok := edges[temp1]; !ok {
			edges[temp1] = []int{temp2}
		} else {
			edges[temp1] = append(edges[temp1], temp2)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	return n, nodes, edges
}

func main() {
	problemWrapper2()
}
