package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-one-row-to-tree/

// Given the root of a binary tree, then value v and depth d, you need to add a row of nodes with value v at the given depth d. The root node is at depth 1.

// The adding rule is: given a positive integer depth d, for each NOT null tree nodes N in depth d-1, create two tree nodes with value v as N's left subtree root and right subtree root. And N's original left subtree should be the left subtree of the new left subtree root, its original right subtree should be the right subtree of the new right subtree root. If depth d is 1 that means there is no depth d-1 at all, then create a tree node with value v as the new root of the whole original tree, and the original tree is the new root's left subtree.

// Note:
// The given d is in range [1, maximum depth of the given tree + 1].
// The given binary tree has at least one tree node.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	switch d {
	case 1:
		newRoot := &TreeNode{Val: v, Left: root}
		return newRoot
	case 2:
		left := &TreeNode{Val: v, Left: root.Left}
		right := &TreeNode{Val: v, Right: root.Right}
		root.Left = left
		root.Right = right
	default:
		if root.Left != nil {
			root.Left = addOneRow(root.Left, v, d-1)
		}
		if root.Right != nil {
			root.Right = addOneRow(root.Right, v, d-1)
		}
	}
	return root
}

func TestAddOneRow(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		v, d   int
		output *TreeNode
	}{
		{Ints2Tree([]int{4, 2, 6, 3, 1, 5}), 1, 2, Ints2Tree([]int{4, 1, 1, 2, null, null, 6, 3, 1, 5})},
		{Ints2Tree([]int{4, 2, null, 3, 1}), 1, 3, Ints2Tree([]int{4, 2, null, 1, 1, 3, null, null, 1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, addOneRow(ts.root, ts.v, ts.d))
		})
	}
}
