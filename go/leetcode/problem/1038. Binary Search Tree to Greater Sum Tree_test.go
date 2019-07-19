package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

// Given the root of a binary search tree with distinct values, modify it so that every node has a new value equal to the sum of the values of the original tree that are greater than or equal to node.val.

// As a reminder, a binary search tree is a tree that satisfies these constraints:

// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	helper1038(root, 0)
	return root
}

func helper1038(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	node.Val += helper1038(node.Right, sum)
	return helper1038(node.Left, node.Val)
}

func TestBstToGst(t *testing.T) {
	tests := []struct {
		input, output *TreeNode
	}{
		{Ints2Tree([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}), Ints2Tree([]int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.Tree2Ints(), bstToGst(ts.input).Tree2Ints())
		})
	}
}
