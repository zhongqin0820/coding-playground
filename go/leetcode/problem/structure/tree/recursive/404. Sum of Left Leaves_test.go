package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-left-leaves/description/

// Find the sum of all left leaves in a given binary tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左子树为空
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}

	// 左叶子节点
	if root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func TestSumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), 24},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sumOfLeftLeaves(ts.root))
		})
	}
}
