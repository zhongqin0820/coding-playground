package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

// Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
func lowestCommonAncestor(root, p, q *TreeNode) int {
	// BST的特性，左 < 中 < 右
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root.Val
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root, p, q *TreeNode
		output     int
	}{
		{Ints2Tree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}), Ints2Tree([]int{2}), Ints2Tree([]int{8}), 6},
		// Explanation: The LCA of nodes 2 and 8 is 6.
		{Ints2Tree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}), Ints2Tree([]int{2}), Ints2Tree([]int{4}), 2},
		// Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, lowestCommonAncestor(ts.root, ts.p, ts.q))
		})
	}
}
