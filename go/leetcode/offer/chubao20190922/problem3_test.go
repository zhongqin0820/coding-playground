package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem3(t *testing.T) {
	tests := []struct {
		A      int
		B      []int
		output int
	}{
		{6, []int{-1, 0, 1, 2, -1, -4}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem3(ts.A, ts.B))
		})
	}
}
