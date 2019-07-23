package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/trim-a-binary-search-tree/

// Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

func TestTrimBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		L, R   int
		output *TreeNode
	}{
		{Ints2Tree([]int{1, 0, 2}), 1, 2, Ints2Tree([]int{1, null, 2})},
		// {Ints2Tree([]int{3, 0, 4, null, 2, null, null, null, null, 1}), 1, 3, Ints2Tree([]int{3, 2, null, 1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, trimBST(ts.root, ts.L, ts.R))
		})
	}
}
