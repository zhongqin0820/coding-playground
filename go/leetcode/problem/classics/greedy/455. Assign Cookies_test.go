package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/assign-cookies/description/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g      []int
		s      []int
		output int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findContentChildren(ts.g, ts.s))
		})
	}
}
