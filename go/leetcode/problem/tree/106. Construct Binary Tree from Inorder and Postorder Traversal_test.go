package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// Given inorder and postorder traversal of a tree, construct the binary tree.

// Note:
// You may assume that duplicates do not exist in the tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return res
	}

	idx := 0
	for i, v := range inorder {
		if v == res.Val {
			idx = i
			break
		}
	}

	res.Left = buildTree(inorder[:idx], postorder[:idx])
	res.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return res
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder, postorder []int
		output              *TreeNode
	}{
		{[]int{}, []int{}, nil},
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, Ints2Tree([]int{3, 9, 20, null, null, 15, 7})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), buildTree(ts.preorder, ts.postorder).PrintTreeNode())
		})
	}
}
