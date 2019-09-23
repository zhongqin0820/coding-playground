package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		A, B   []int
		output []int
	}{
		{[]int{1, 3, 7, 4, 10}, []int{2, 1, 5, 8, 9}, []int{1, 3, 7, 9, 10}},
		{[]int{20, 1, 23}, []int{50, 26, 7}, []int{}},
		{[]int{1}, []int{2, 3}, []int{1}},
		{[]int{}, []int{2, 3}, []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem1(ts.A, ts.B))
		})
	}
}
