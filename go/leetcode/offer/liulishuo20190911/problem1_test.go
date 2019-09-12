package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		n, a, b int
		c, d    []int
		output  int
	}{
		{3, 1, 2, []int{13, 4, 10}, []int{19, 9, 20}, 38},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem1(ts.n, ts.a, ts.b, ts.c, ts.d))
		})
	}
}
