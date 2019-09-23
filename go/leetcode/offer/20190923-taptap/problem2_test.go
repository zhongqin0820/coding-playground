package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2(t *testing.T) {
	tests := []struct {
		A, B   []int
		output []int
	}{
		{[]int{}, []int{}, []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem2(ts.A, ts.B))
		})
	}
}

// 2
// 61 80
// 0 1

// 80
// //
// 3
// 30 13 41
// 1 0
// 1 2

// 54
// //
// 4
// 33 72 19 54
// 0 2
// 1 3
// 1 2

// 145
// //
// 12
// 29 29 48 38 6 38 28 24 6 62 38 84
// 9 8
// 10 2
// 2 3
// 10 6
// 5 1
// 4 5
// 6 7
// 8 5
// 9 6
// 4 0
// 11 1

// 401
