package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-width-of-binary-tree/

// Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.

// The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.

// Note: Answer will in the range of 32-bit signed integer.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	// ref: https://leetcode.com/problems/maximum-width-of-binary-tree/discuss/269657/Golang-InOrder-Traversal-with-comments-O(n)-complexity
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	widths := make(map[int][]int)

	var in func(*TreeNode, int, int)
	in = func(node *TreeNode, level, index int) {
		if node == nil {
			return
		}
		//
		in(node.Left, level+1, (index*2)-1)
		//
		if _, ok := widths[level]; !ok {
			widths[level] = []int{}
		}
		widths[level] = append(widths[level], index)
		//
		in(node.Right, level+1, index*2)
	}

	in(root, 0, 1)
	//
	max, width := 1, 0
	for _, row := range widths {
		if n := len(row); 1 < n {
			width = row[n-1] - row[0] + 1
		}
		if max < width {
			max = width
		}
	}
	return max
}

func TestWidthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output int
	}{
		{Ints2Tree([]int{1, 3, 2, 5, 3, null, 9}), 4},
		{Ints2Tree([]int{1, 3, null, 5, 3}), 2},
		{Ints2Tree([]int{1, 3, 2, 5}), 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, widthOfBinaryTree(ts.root))
		})
	}
}
