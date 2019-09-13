package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		A      [][]int
		n, m   int
		output int
	}{
		{[][]int{{0, 1, 0, 0, 0}, {1, 1, 1, 0, 0}, {1, 1, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 1, 0, 1}}, 5, 5, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem1(ts.A, ts.n, ts.m))
		})
	}
}
