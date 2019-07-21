package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/flip-equivalent-binary-trees/

// For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.

// A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

// Write a function that determines whether two binary trees are flip equivalent.  The trees are given by root nodes root1 and root2.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	return root1.Val == root2.Val && ((flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || (flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)))
}

func TestFlipEquiv(t *testing.T) {
	tests := []struct {
		root1  *TreeNode
		root2  *TreeNode
		output bool
	}{
		{Ints2Tree([]int{1, 2, 3, 4, 5, 6, null, null, null, 7, 8}), Ints2Tree([]int{1, 3, 2, null, 6, 4, 5, null, null, null, null, 8, 7}), true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, flipEquiv(ts.root1, ts.root2))
		})
	}
}
