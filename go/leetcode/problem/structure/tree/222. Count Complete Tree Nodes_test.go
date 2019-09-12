package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-complete-tree-nodes/

// Given a complete binary tree, count the number of nodes.

// Note:

// Definition of a complete binary tree from Wikipedia:
// In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	count := 0

	var pre func(*TreeNode)
	pre = func(root *TreeNode) {
		if root == nil {
			return
		}
		count++
		pre(root.Left)
		pre(root.Right)
	}
	pre(root)

	return count
}

func TestCountNodes(t *testing.T) {
	tests := []struct {
		intput *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, 2, 3, 4, 5, 6}), 6},
		{Ints2Tree([]int{1}), 1},
		{Ints2Tree([]int{1, 2, 3}), 3},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, countNodes(ts.intput))
		})
	}
}
