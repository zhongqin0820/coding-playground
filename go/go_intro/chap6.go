package main

import (
	"fmt"
	"log"
	"time"
)

func exp1() {
	log.Println(min(1, 2, 4, -1))
	log.Println(min([]int{1, 2, 3, -5, 6}...))
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// defer is implement at the very end of a function
func exp2() {
	fmt.Println("start")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

func exp3() {
	start := time.Now()
	recursivePrintI(10)
	end := time.Now()
	duration := end.Sub(start)
	log.Println(duration)
}

func recursivePrintI(i int) {
	if i == 0 {
		return
	}
	log.Println(i)
	i--
	recursivePrintI(i)
}

// lambda function and defer keywords
func exp4() {
	res := func() (ret int) {
		defer func() {
			ret += 4
		}()
		return 1
	}()
	log.Println(res)
}

func main() {
	// exp1()
	// exp2()
	exp3()
	// exp4()
	fmt.Println("...")
}
