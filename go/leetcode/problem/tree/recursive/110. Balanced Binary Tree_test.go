package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/balanced-binary-tree/

// Given a binary tree, determine if it is height-balanced.

// For this problem, a height-balanced binary tree is defined as:

// a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	var dfs func(*TreeNode) (int, bool)
	dfs = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		l, lB := dfs(root.Left)
		r, rB := dfs(root.Right)
		if lB && rB && -1 <= l-r && l-r <= 1 {
			return max(l, r) + 1, true
		}
		return 0, false
	}

	_, b := dfs(root)
	return b
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output bool
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), true},
		{Ints2Tree([]int{1, 2, 2, 3, 3, null, null, 4, 4}), false},
		{Ints2Tree([]int{}), true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isBalanced(ts.root))
		})
	}
}
