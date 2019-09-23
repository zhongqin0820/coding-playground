package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		A, B   []int
		output []int
	}{
		{[]int{}, []int{}, []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, problem1(ts.A, ts.B))
		})
	}
}

// abcdefghijklmnopqrstuvwxyz
// 3
// cba
// def
// abc

// abc
// cba
// def

// zyxwvutsrqponmlkjihgfedcba
// 3
// aaa
// bbb
// ccc

// ccc
// bbb
// aaa

// qwertyuiopasdfghjklzxcvbnm
// 2
// wert
// werty

// wert
// werty

// qazxswedcvfrtgbnhyujmkilop
// 4
// oooooo
// xwcvf
// xartnhg
// xatghs

// xartnhg
// xatghs
// xwcvf
// oooooo
