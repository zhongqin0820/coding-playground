package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-paths/description/

// Given a binary tree, return all root-to-leaf paths.
// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0, 16)

	var dfs func(string, *TreeNode)
	dfs = func(pre string, node *TreeNode) {
		if pre == "" {
			pre = strconv.Itoa(node.Val)
		} else {
			pre += "->" + strconv.Itoa(node.Val)
		}

		if node.Left != nil {
			dfs(pre, node.Left)
		}
		if node.Right != nil {
			dfs(pre, node.Right)
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, pre)
		}
	}

	dfs("", root)
	return res
}

func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []string
	}{
		{Ints2Tree([]int{1, 2, 3, null, 5}), []string{"1->2->5", "1->3"}},
		// Explanation: All root-to-leaf paths are: 1->2->5, 1->3
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, binaryTreePaths(ts.root))
		})
	}
}
