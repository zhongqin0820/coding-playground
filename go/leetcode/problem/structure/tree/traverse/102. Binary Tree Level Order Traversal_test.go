package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-tree-level-order-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := [][]int{}
	for n := len(queue); n != 0; n = len(queue) {
		// 内循环遍历当前层
		temp := make([]int, 0, 2048)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output [][]int
	}{
		{Ints2Tree([]int{3, 9, 20, null, null, 15, 7}), [][]int{
			[]int{3},
			[]int{9, 20},
			[]int{15, 7},
		}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, levelOrder(ts.root))
		})
	}
}
