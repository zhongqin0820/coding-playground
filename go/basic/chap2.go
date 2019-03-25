package main

import (
	"fmt"
)

type Q struct {
}

// Q6
func (q *Q) q6() float64 {
	s := []float64{1.1, 1.2, 1.3}
	var r float64
	for _, v := range s {
		r += v
	}
	m := r / float64(len(s))
	return m
}

// Q7
func (q *Q) q7(a, b int) []int {
	if a <= b {
		return []int{a, b}
	} else {
		return []int{b, a}
	}
}

// type mStack struct {
// 	v []int
// }

// // Q9
// func (q *Q) q9() {

// }
func main() {
	q := Q{}
	// m := q.q6()
	// fmt.Println(m)
	// s := q.q7(2, 7)
	// fmt.Println(s)
}
