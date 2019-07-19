package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

// Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

// Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val <= val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func TestInsertIntoBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		val    int
		output *TreeNode
	}{
		{Ints2Tree([]int{4, 2, 7, 1, 3}), 5, Ints2Tree([]int{4, 2, 7, 1, 3, 5})},
		{Ints2Tree(nil), 4, Ints2Tree([]int{4})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), insertIntoBST(ts.root, ts.val).PrintTreeNode())
		})
	}
}
