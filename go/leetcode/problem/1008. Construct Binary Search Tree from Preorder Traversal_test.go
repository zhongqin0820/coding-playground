package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

// Return the root node of a binary search tree that matches the given preorder traversal.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	i, n := 0, len(preorder)
	var helper func(int) *TreeNode
	helper = func(bound int) *TreeNode {
		if i == n || preorder[i] > bound {
			return nil
		}
		root := &TreeNode{Val: preorder[i]}
		i++
		root.Left = helper(root.Val)
		root.Right = helper(bound)
		return root
	}
	return helper(1 << 32)
}

func TestBstFromPreorder(t *testing.T) {
	tests := []struct {
		input  []int
		output *TreeNode
	}{
		{[]int{8, 5, 1, 7, 10, 12}, Ints2Tree([]int{8, 5, 10, 1, 7, null, 12})},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), bstFromPreorder(ts.input).PrintTreeNode())
		})
	}
}
