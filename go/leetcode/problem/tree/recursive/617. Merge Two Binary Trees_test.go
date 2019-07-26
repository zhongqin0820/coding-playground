package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-binary-trees/description/

// Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.
// You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil && t2 != nil {
		return t2
	}
	if t1 != nil && t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func TestMergeTrees(t *testing.T) {
	tests := []struct {
		t1, t2 *TreeNode
		output *TreeNode
	}{
		{Ints2Tree([]int{1, 3, 2, 5}), Ints2Tree([]int{2, 1, 3, null, 4, null, 7}), Ints2Tree([]int{3, 4, 5, 5, 4, null, 7})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, mergeTrees(ts.t1, ts.t2))
		})
	}
}
