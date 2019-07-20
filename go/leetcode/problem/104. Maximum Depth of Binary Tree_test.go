package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

// Given a binary tree, find its maximum depth.

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper104(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func helper104(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxDepth(ts.input))
		})
	}
}
