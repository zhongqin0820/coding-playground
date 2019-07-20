package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-depth-of-binary-tree/

// Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + helper111(minDepth(root.Left), minDepth(root.Right))
	}
}

func helper111(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), 2},
		{Ints2Tree([]int{}), 0},
		{Ints2Tree([]int{1, 2, 3}), 2},
		{Ints2Tree([]int{1, 2}), 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minDepth(ts.input))
		})
	}
}
