package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/validate-binary-tree-nodes/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// 统计树的出入度
	var in = make([]int, n)
	var out = make([]int, n)
	for i, l := range leftChild {
		if l == -1 {
			continue
		}
		in[l]++
		out[i]++
	}
	for i, r := range rightChild {
		if r == -1 {
			continue
		}
		in[r]++
		out[i]++
	}
	zero := 0
	for _, cnt := range in {
		if cnt == 0 {
			zero++
		}
		if cnt > 1 {
			return false
		}
	}
	if zero != 1 {
		return false
	}
	for _, cnt := range out {
		if cnt > 2 {
			return false
		}
	}
	return true
}

func TestValidateBinaryTreeNodes(t *testing.T) {
	tests := []struct {
		n          int
		leftChild  []int
		rightChild []int
		output     bool
	}{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{2, []int{1, 0}, []int{-1, -1}, false},
		{6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, validateBinaryTreeNodes(ts.n, ts.leftChild, ts.rightChild))
		})
	}
}
