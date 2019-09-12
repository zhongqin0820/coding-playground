package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/construct-string-from-binary-tree/

// You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.

// The null node needs to be represented by empty parenthesis pair "()". And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	var pre func(*TreeNode) string
	pre = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		if node.Left == nil && node.Right == nil {
			return strconv.Itoa(node.Val)
		}
		if node.Right == nil {
			return strconv.Itoa(node.Val) + "(" + pre(node.Left) + ")"
		}
		return strconv.Itoa(node.Val) + "(" + pre(node.Left) + ")" + "(" + pre(node.Right) + ")"
	}
	return pre(t)
}
func TestTree2str(t *testing.T) {
	tests := []struct {
		t      *TreeNode
		output string
	}{
		{Ints2Tree([]int{1, 2, 3, 4}), "1(2(4))(3)"},
		// Explanation: Originallay it needs to be "1(2(4)())(3()())",
		// but you need to omit all the unnecessary empty parenthesis pairs.
		// And it will be "1(2(4))(3)".
		{Ints2Tree([]int{1, 2, 3, null, 4}), "1(2()(4))(3)"},
		// Explanation: Almost the same as the first example,
		// except we can't omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, tree2str(ts.t))
		})
	}
}
