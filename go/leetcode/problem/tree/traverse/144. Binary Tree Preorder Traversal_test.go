package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-preorder-traversal/description/

// Given a binary tree, return the preorder traversal of its nodes' values.

// Example:

// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [1,2,3]
// Follow up: Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	// 非递归遍历
	stack, res := []*TreeNode{}, []int{}
	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil { //遍历左侧
			if cur.Right != nil { //右侧节点压栈
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
		} else { // cur.Left == nil
			if cur.Right != nil {
				cur = cur.Right
			} else { // 当前节点为叶子
				if len(stack) == 0 { //栈中没有节点，退出
					break
				}
				//从栈中取出节点
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
	}

	return res
}

func preorderTraversalAdv(root *TreeNode) []int {
	// 递归遍历版本
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversalAdv(root.Left)...)
	res = append(res, preorderTraversalAdv(root.Right)...)
	return res
}

func preorderTraversalAdvII(root *TreeNode) []int {
	stack := []*TreeNode{root}
	res := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}

		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, null, 2, 3}), []int{1, 2, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, preorderTraversal(ts.root))
			ast.Equal(ts.output, preorderTraversalAdv(ts.root))
			ast.Equal(ts.output, preorderTraversalAdvII(ts.root))
		})
	}
}
