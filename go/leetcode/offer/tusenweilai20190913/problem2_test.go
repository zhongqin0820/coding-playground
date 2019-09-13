package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	tests := []struct {
		A, B   []string
		n      int
		output []int
	}{
		{[]string{"134503"}, []string{"834703"}, 1, []int{2}},
		{[]string{"000001"}, []string{"000017"}, 1, []int{-1}},
		{[]string{"134503", "134503"}, []string{"834703", "1034703"}, 2, []int{2, -1}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem2(ts.A, ts.B, ts.n))
		})
	}
}
