package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/recover-binary-search-tree/

// Two elements of a binary search tree (BST) are swapped by mistake.

// Recover the tree without changing its structure.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) *TreeNode {
	var first, second, prev *TreeNode
	// first记录第一次降序的最大值
	// second记录第二次降序的最小值
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev
			}
			if first != nil {
				second = root
			}
		}
		prev = root

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)
	first.Val, second.Val = second.Val, first.Val
	return root
}

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output *TreeNode
	}{
		{Ints2Tree([]int{1, 3, null, null, 2}), Ints2Tree([]int{3, 1, null, null, 2})},
		{Ints2Tree([]int{3, 1, 4, null, null, 2}), Ints2Tree([]int{2, 1, 4, null, null, 3})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), recoverTree(ts.root).PrintTreeNode())
		})
	}
}
