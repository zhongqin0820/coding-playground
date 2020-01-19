package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-172/problems/delete-leaves-with-a-given-value/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == root.Right && root.Val == target {
		return nil
	}
	return root
}

func TestRemoveLeafNodes(t *testing.T) {
	// Constraints:
	//     1 <= target <= 1000
	//     Each tree has at most 3000 nodes.
	//     Each node's value is between [1, 1000].
	tests := []struct {
		root   *TreeNode
		target int
		output *TreeNode
	}{
		{Ints2Tree([]int{1, 2, 3, 2, null, 2, 4}), 2, Ints2Tree([]int{1, null, 3, null, 4})},
		{Ints2Tree([]int{1, 3, 3, 3, 2}), 3, Ints2Tree([]int{1, 3, null, null, 2})},
		{Ints2Tree([]int{1, 2, null, 2, null, 2}), 2, Ints2Tree([]int{1})},
		{Ints2Tree([]int{1, 1, 1}), 1, Ints2Tree([]int{})},
		{Ints2Tree([]int{1, 2, 3}), 1, Ints2Tree([]int{1, 2, 3})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, removeLeafNodes(ts.root, ts.target))
		})
	}
}
