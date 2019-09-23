package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	// replace the case struct
	tests := []struct {
		A, B   int
		output int
	}{
		{151, 160, 8},
		{152, 160, 5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			// replace the problem
			ast.Equal(ts.output, problem(ts.A, ts.B))
		})
	}
}
