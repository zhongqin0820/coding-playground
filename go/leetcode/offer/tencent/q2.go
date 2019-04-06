package main

import (
	"fmt"
)

func main() {
	N := 0
	s := ""
	fmt.Scan(&N)
	fmt.Scan(&s)
	if len(s)%2 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
