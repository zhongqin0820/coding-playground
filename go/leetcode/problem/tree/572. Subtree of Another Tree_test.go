package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subtree-of-another-tree/

// Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper572(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return helper572(s.Left, t.Left) && helper572(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if helper572(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		s, t   *TreeNode
		output bool
	}{
		{Ints2Tree([]int{3, 4, 5, 1, 2}), Ints2Tree([]int{4, 1, 2}), true},
		{Ints2Tree([]int{3, 4, 5, 1, null, 2}), Ints2Tree([]int{3, 1, 2}), false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isSubtree(ts.s, ts.t))
		})
	}
}
