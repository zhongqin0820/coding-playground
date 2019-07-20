package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-root-to-leaf-numbers/

// Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

// An example is the root-to-leaf path 1->2->3 which represents the number 123.

// Find the total sum of all root-to-leaf numbers.

// Note: A leaf is a node with no children.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, temp int) {
		if node == nil {
			return
		}

		temp = temp*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += temp
			return
		}

		dfs(node.Left, temp)
		dfs(node.Right, temp)
	}

	dfs(root, 0)
	return res
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1}), 1},
		{Ints2Tree([]int{1, 2, 3}), 25},
		// Explanation:
		//         The root-to-leaf path 1->2 represents the number 12.
		//         The root-to-leaf path 1->3 represents the number 13.
		//         Therefore, sum = 12 + 13 = 25.
		{Ints2Tree([]int{4, 9, 0, 5, 1}), 1026},
		// Explanation:
		//         The root-to-leaf path 4->9->5 represents the number 495.
		//         The root-to-leaf path 4->9->1 represents the number 491.
		//         The root-to-leaf path 4->0 represents the number 40.
		//         Therefore, sum = 495 + 491 + 40 = 1026.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sumNumbers(ts.root))
		})
	}
}
