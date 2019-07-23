package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/average-of-levels-in-binary-tree/

// Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

// Note:
// The range of node's value is in the range of 32-bit signed integer.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	last, nlast := root, root
	cur := []int{}
	res := []float64{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		cur = append(cur, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
			nlast = node.Left
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nlast = node.Right
		}
		if last == node {
			curRes := 0
			for _, v := range cur {
				curRes += v
			}
			res = append(res, float64(curRes)/float64(len(cur)))
			cur = []int{}
			last = nlast
		}
	}
	return res
}

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []float64
	}{
		{Ints2Tree([]int{3, 9, 20, 15, 7}), []float64{3, 14.5, 11}},
		// Explanation:
		// The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, averageOfLevels(ts.root))
		})
	}
}
