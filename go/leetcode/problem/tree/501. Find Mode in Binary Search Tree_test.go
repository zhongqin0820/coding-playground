package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-mode-in-binary-search-tree/

// Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

// Assume a BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than or equal to the node's key.
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
// Both the left and right subtrees must also be binary search trees.

// Note: If a tree has more than one mode, you can return them in any order.

// Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	r := map[int]int{}
	var pre func(*TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		r[node.Val]++
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)

	max := -1 << 63
	res := []int{}
	for k, v := range r {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, k)
		}
	}
	return res
}

func TestFindMode(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, null, 2, 2}), []int{2}},
		{Ints2Tree([]int{1, null, 2, 2, 2, null, null, 3, 3, 3}), []int{2, 3}},
		{Ints2Tree([]int{1, null, 2, 2, 2, null, null, 3, 3, 3, 3}), []int{3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMode(ts.root))
		})
	}
}
