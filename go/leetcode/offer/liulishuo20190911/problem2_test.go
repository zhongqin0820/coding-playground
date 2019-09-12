package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	tests := []struct {
		maze       [][]int
		start, end point
		output     int
	}{
		{
			[][]int{
				[]int{2, 0, 0, 1},
				[]int{0, 1, 0, 1},
				[]int{1, 0, 0, 0},
				[]int{3, 0, 1, 0},
			},
			point{0, 0}, point{3, 0},
			7,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem2(ts.maze, ts.start, ts.end))
		})
	}
}
