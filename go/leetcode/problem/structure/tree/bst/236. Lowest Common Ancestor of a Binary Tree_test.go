package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestorII(root.Left, p, q)  //p,q在root.Left中？
	r := lowestCommonAncestorII(root.Right, p, q) //p,q在root.Right中？

	if l != nil && r != nil {
		// p,q分别在两侧
		return root
	}

	if l == nil {
		// p,q在右侧
		return r
	}
	// p,q在左侧
	return l
}

func TestLowestCommonAncestorII(t *testing.T) {
	tests := []struct {
		root, p, q *TreeNode
		output     *TreeNode
	}{
		{Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}), Ints2Tree([]int{5}), Ints2Tree([]int{1}), Ints2Tree([]int{3})},
		// Explanation: The LCA of nodes 5 and 1 is 3.
		{Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}), Ints2Tree([]int{5}), Ints2Tree([]int{4}), Ints2Tree([]int{5})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			res := lowestCommonAncestorII(ts.root, ts.p, ts.q)
			if res != nil {
				ast.Equal(ts.output.Val, res.Val)
			}
		})
	}
}
