package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	tests := []struct {
		a      string
		output bool
	}{
		{"ababa;ststs", true},
		{"ababa;stsas", false},
		{";", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem2(ts.a))
		})
	}
}
