package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	// replace the case struct
	tests := []struct {
		A      string
		output string
	}{
		{"8 1 2 3 4 5 6 A", "9 1 2 3 4 5 6 12 34"},
		{"A 1 2 3 4 5 6 A 7 B", "C 1 2 3 4 5 6 12 34 7 AB CD"},
		{"13 1 2 3 4 5 6 A 7 B 9 1 4 5 9 1 A B", "16 1 2 3 4 5 6 12 34 7 AB CD 9 1 4 5 9 1 12 34 AB CD"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			// replace the problem
			ast.Equal(ts.output, problem1(ts.A))
		})
	}
}
