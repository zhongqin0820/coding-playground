package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-distance-between-bst-nodes/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	res, pre, null := 1<<63-1, 1>>63, 1>>63
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}

		if pre != null {
			res = helper783Min(res, node.Val-pre)
		}
		pre = node.Val

		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return res
}

func helper783Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinDiffInBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{4, 2, 6, 1, 3, null, null}), 1},
		{Ints2Tree([]int{90, 69, null, 49, 89, null, 52, null, null, null, null}), 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minDiffInBST(ts.root))
		})
	}
}
