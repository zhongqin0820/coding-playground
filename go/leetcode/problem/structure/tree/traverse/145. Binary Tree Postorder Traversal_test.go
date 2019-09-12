package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given a binary tree, return the postorder traversal of its nodes' values.

// Example:

// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// Output: [3,2,1]
// Follow up: Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	// 非递归
	stack := []*TreeNode{root}
	res := []int{}
	// root->right->left
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	// reverse res
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

func postorderTraversalAdv(root *TreeNode) []int {
	// 递归
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, postorderTraversalAdv(root.Left)...)
	res = append(res, postorderTraversalAdv(root.Right)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversalAdvII(root *TreeNode) []int {
	// 牛客算法专题的解析方法其实和第一种方法思路是一样的
	s1, s2 := []*TreeNode{root}, []int{}
	for len(s1) > 0 {
		node := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		if node == nil {
			continue
		}
		s2 = append(s2, node.Val)
		s1 = append(s1, node.Left)
		s1 = append(s1, node.Right)
	}
	// s2 是结果栈，出栈顺序即为结果，将其翻转即可
	// reverse s2
	for i, n := 0, len(s2); i < n/2; i++ {
		s2[i], s2[n-i-1] = s2[n-i-1], s2[i]
	}
	return s2
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, null, 2, 3}), []int{3, 2, 1}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, postorderTraversal(ts.root))
			ast.Equal(ts.output, postorderTraversalAdv(ts.root))
			ast.Equal(ts.output, postorderTraversalAdvII(ts.root))
		})
	}
}
