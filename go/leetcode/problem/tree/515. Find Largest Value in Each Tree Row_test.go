package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

// You need to find the largest value in each row of a binary tree.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	cur, res := -1<<63, []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	last, nlast := root, root
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val > cur {
			cur = node.Val
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nlast = node.Left
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nlast = node.Right
		}
		if last == node {
			res = append(res, cur)
			cur = -1 << 63
			last = nlast
		}
	}
	return res
}

func TestLargestValues(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{1, 3, 2, 5, 3, 9}), []int{1, 3, 9}},
		{Ints2Tree([]int{}), []int{}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, largestValues(ts.root))
		})
	}
}
