package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput1(t *testing.T) {
	tests := []struct {
		input  string
		output []int
	}{
		{"1 2 3 4 5 6 7", []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, Input1(ts.input))
		})
	}
}

func TestInput2(t *testing.T) {
	tests := []struct {
		input   string
		output1 []int
		output2 []int
	}{
		{`1 2
1 2
3 4
5 6
7 8
`, []int{1, 1, 3, 5, 7}, []int{2, 2, 4, 6, 8}},
		{`1 2
1 2
3 4
5 6
7 8



`, []int{1, 1, 3, 5, 7}, []int{2, 2, 4, 6, 8}},
		{`1 2



`, []int{1}, []int{2}},
		{`



`, []int{}, []int{}},
		{`1 2
3 4
5 

`, []int{1, 3}, []int{2, 4}},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			output1, output2 := Input2(ts.input)
			ast.Equal(ts.output1, output1)
			ast.Equal(ts.output2, output2)
		})
	}
}
