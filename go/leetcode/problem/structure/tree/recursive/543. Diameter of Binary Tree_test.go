package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/diameter-of-binary-tree/description/

// Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

// Note: The length of path between two nodes is represented by the number of edges between them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	dfs(root)
	return res
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, 2, 3, 4, 5}), 3},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, diameterOfBinaryTree(ts.root))
		})
	}
}
