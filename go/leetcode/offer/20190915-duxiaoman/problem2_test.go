package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	tests := []struct {
		N, A, B, C int
		a          []int
		output     int
	}{
		{7, 1, 1, 1, []int{3, 6, 4, 3, 4, 5, 6}, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem2(ts.N, ts.A, ts.B, ts.C, ts.a))
		})
	}
}
