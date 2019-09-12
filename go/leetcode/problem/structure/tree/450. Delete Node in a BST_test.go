package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/delete-node-in-a-bst/

// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

// Basically, the deletion can be divided into two stages:
// 1. Search for a node to remove.
// 2. If the node is found, delete the node.
// Note: Time complexity should be O(height of tree).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key == root.Val {
		if root.Left == nil && root == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// find the successor
			// the max val on the left part of the node
			// delete recursively
			max := helper450(root.Left)
			root.Val = max
			root.Left = deleteNode(root.Left, max)
		}
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func helper450(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return helper450(root.Right)
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		key    int
		output *TreeNode
	}{
		{Ints2Tree([]int{5, 3, 6, 2, 4, null, 7}), 3, Ints2Tree([]int{5, 2, 6, null, 4, null, 7})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, deleteNode(ts.root, ts.key))
		})
	}
}
