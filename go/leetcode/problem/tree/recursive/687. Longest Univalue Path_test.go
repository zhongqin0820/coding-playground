package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

// The length of path between two nodes is represented by the number of edges between them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper687(root, &maxLen)
	return maxLen
}

// 返回从 root 出发拥有相同 Val 值的线路上的 edge 数量
func helper687(root *TreeNode, maxLen *int) int {
	if root == nil {
		return 0
	}
	l := helper687(root.Left, maxLen)
	r := helper687(root.Right, maxLen)

	res := 0
	// 左侧单边最长
	if root.Left != nil && root.Val == root.Left.Val {
		*maxLen = helper687Max(*maxLen, l+1)
		res = helper687Max(res, l+1)
	}

	// 右侧单边最长
	if root.Right != nil && root.Val == root.Right.Val {
		*maxLen = helper687Max(*maxLen, r+1)
		res = helper687Max(res, r+1)
	}

	// 通过根节点
	if root.Left != nil && root.Val == root.Left.Val && root.Right != nil && root.Val == root.Right.Val {
		*maxLen = helper687Max(*maxLen, l+r+2)
	}

	return res
}

func helper687Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestLongestUnivaluePath(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{5, 4, 5, 1, 1, 5}), 2},
		{Ints2Tree([]int{1, 4, 5, 4, 4, 5}), 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, longestUnivaluePath(ts.root))
		})
	}
}
