package main

import (
	"fmt"
)

func main() {
	N := 0
	d := []int{}
	p := []int{}
	fmt.Scan(&N)
	var temp int
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		d = append(d, temp)
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		p = append(p, temp)
	}
	curD := d[0]
	curP := p[0]
	// fmt.Println()
	// fmt.Println()
	// fmt.Println(curD, curP)
	curP = helper(d, p, curD, curP, 1)
	fmt.Println(curP)
}

func helper(d, p []int, cD, cP, i int) int {
	if len(d) == i {
		// fmt.Println("len D=", len(d), " i=", i)
		return cP
	}
	tempD := d[i]
	tempP := p[i]
	// fmt.Println(cD, cP, tempD, tempP)
	if tempD > cD {
		cD = cD + tempD
		cP = cP + tempP
	}
	return min(helper(d, p, cD, cP, i+1), helper(d, p, cD+tempD, cP+tempP, i+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
