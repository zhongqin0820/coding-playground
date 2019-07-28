package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem4(t *testing.T) {
	tests := []struct {
		n      int
		L, W   []int
		output int
	}{
		{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 10}, 9},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem4(ts.n, ts.L, ts.W))
		})
	}
}
