package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem3(t *testing.T) {
	tests := []struct {
		A      string
		B      int
		C      [][]string
		output int
	}{
		{"Jack", 3, [][]string{[]string{"Jack", "Tom", "Anny", "Lucy"}, []string{"Tom", "Danny"}, []string{"Jack", "Lily"}}, 6},
		{"ABC", 3, [][]string{[]string{"Jack", "Tom", "Anny", "Lucy"}, []string{"Tom", "Danny"}, []string{"Jack", "Lily"}}, 0},
		{"Mike", 3, [][]string{[]string{"Jack", "Tom", "Anny", "Lucy"}, []string{"Mike", "Danny"}, []string{"Jack", "Lily"}}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem3(ts.A, ts.B, ts.C))
		})
	}
}
