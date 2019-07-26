package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-sum/

// Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		sum    int
		output bool
	}{
		{Ints2Tree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}), 22, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, hasPathSum(ts.root, ts.sum))
		})
	}
}
