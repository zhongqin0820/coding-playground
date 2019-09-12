package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

// Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	// inorder traverse 将BST遍历得到生序序列匹配另一个值target = k-node.Val
	var in func(*TreeNode, int) bool
	in = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}
		if node.Val == target {
			return true
		}
		if node.Val < target {
			return in(node.Right, target)
		}
		return in(node.Left, target)
	}
	// helper
	var helper func(*TreeNode, *TreeNode, int) bool
	helper = func(node, searchNode *TreeNode, target int) bool {
		if node == nil {
			return false
		}
		// node.Val*2 != target indicates two different nodes
		return (node.Val*2 != target && in(searchNode, k-node.Val)) || helper(node.Left, searchNode, k) || helper(node.Right, searchNode, k)
	}

	return helper(root, root, k)
}

func TestFindTarget(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		k      int
		output bool
	}{
		{Ints2Tree([]int{5, 3, 6, 2, 4, null, 7}), 9, true},
		{Ints2Tree([]int{5, 3, 6, 2, 4, null, 7}), 28, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findTarget(ts.root, ts.k))
		})
	}
}
