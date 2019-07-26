package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/invert-binary-tree/description/

// Invert a binary tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTreeAdv(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output *TreeNode
	}{
		{Ints2Tree([]int{4, 2, 7, 1, 3, 6, 9}), Ints2Tree([]int{4, 7, 2, 9, 6, 3, 1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), invertTree(ts.root).PrintTreeNode())
			// ast.Equal(ts.output.PrintTreeNode(), invertTreeAdv(ts.root).PrintTreeNode())
		})
	}
}
