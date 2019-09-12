package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

// Given a binary tree, return the inorder traversal of its nodes' values.

// Example:

// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [1,3,2]
// Follow up: Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for cur := root; cur != nil || len(stack) > 0; {
		for cur != nil { // 压入当前节点左边界
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 左边界弹出节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		// 压入弹出节点的右边界
		cur = node.Right
	}
	return res
}

func inorderTraversalAdv(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, null, 2, 3}), []int{1, 3, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, inorderTraversal(ts.root))
			ast.Equal(ts.output, inorderTraversalAdv(ts.root))
		})
	}
}
