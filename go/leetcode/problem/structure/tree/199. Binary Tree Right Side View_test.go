package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-right-side-view/

// Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{root}
	last, nlast := root, root
	res := []int{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nlast = node.Left
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nlast = node.Right
		}
		if node == last {
			res = append(res, last.Val)
			last = nlast
		}
	}
	return res
}

func TestRightSideView(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, 2, 3, null, 5, null, 4}), []int{1, 3, 4}},
		{Ints2Tree([]int{}), []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, rightSideView(ts.input))
		})
	}
}
