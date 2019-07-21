package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/symmetric-tree/

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper101(root.Left, root.Right)
}

func helper101(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	return helper101(left.Left, right.Right) && helper101(left.Right, right.Left) && left.Val == right.Val
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output bool
	}{
		{Ints2Tree([]int{}), true},
		{Ints2Tree([]int{1, 2, 2, 3, 4, 4, 3}), true},
		{Ints2Tree([]int{1, 2, 2, null, 3, null, 3}), false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isSymmetric(ts.root))
		})
	}
}
