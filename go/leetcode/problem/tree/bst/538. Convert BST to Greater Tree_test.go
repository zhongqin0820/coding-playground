package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-bst-to-greater-tree/description/

// Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var in func(*TreeNode)
	cur := 0 //保存当前（右->中->左）遍历得到的结果
	in = func(node *TreeNode) {
		if node == nil {
			return
		}
		in(node.Right)
		cur += node.Val
		node.Val = cur
		in(node.Left)
	}
	in(root)
	return root
}

func TestConvertBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output *TreeNode
	}{
		{Ints2Tree([]int{5, 2, 13}), Ints2Tree([]int{18, 20, 13})},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintTreeNode(), convertBST(ts.root).PrintTreeNode())
		})
	}
}
