package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem3(t *testing.T) {
	tests := []struct {
		n, m    int
		P, A, B []int
		output  []int
	}{
		{5, 6, []int{1, 2, 1, 1, 1}, []int{1, 1, 1, 2, 3, 4}, []int{2, 3, 4, 5, 5, 5}, []int{1, 3, 4, 2, 5}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem3(ts.n, ts.m, ts.P, ts.A, ts.B))
		})
	}
}
