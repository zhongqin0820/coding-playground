package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		x, y, n int
		xs, ys  []int
		output  int
	}{
		{2, 0, 3, []int{1, 1, 1}, []int{0, 1, -1}, 6},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem1(ts.x, ts.y, ts.n, ts.xs, ts.ys))
		})
	}
}
