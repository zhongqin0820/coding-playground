package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem3(t *testing.T) {
	tests := []struct {
		A      []int
		output int
	}{
		{[]int{2, 4, -2, 5, -6}, 9},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem3(ts.A))
		})
	}
}
