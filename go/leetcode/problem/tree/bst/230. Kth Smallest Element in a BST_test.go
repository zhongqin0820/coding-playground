package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

// Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

// Note:
// You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	// in-order traverse
	var in func(*TreeNode)
	in = func(node *TreeNode) {
		if node == nil {
			return
		}
		in(node.Left)
		if k--; k == 0 {
			res = node.Val
		}
		in(node.Right)
	}
	in(root)
	return res
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		k      int
		output int
	}{
		{Ints2Tree([]int{3, 1, 4, null, 2}), 1, 1},
		{Ints2Tree([]int{5, 3, 6, 2, 4, null, null, 1}), 3, 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, kthSmallest(ts.root, ts.k))
		})
	}
}
