package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-duplicate-subtrees/

// Given a binary tree, return all duplicate subtrees. For each kind of duplicate subtrees, you only need to return the root node of any one of them.

// Two trees are duplicate if they have the same structure with same node values.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int, 1024)
	res := make([]*TreeNode, 0, 1024)
	var pos func(*TreeNode) string
	pos = func(node *TreeNode) string {
		if node == nil {
			return "*"
		}
		l := pos(node.Left)
		r := pos(node.Right)
		key := strconv.Itoa(node.Val) + l + r

		count[key]++
		if count[key] == 2 {
			res = append(res, node)
		}
		return key
	}

	pos(root)
	return res
}

func TestFindDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []*TreeNode
	}{
		{Ints2Tree([]int{1, 2, 3, 4, null, 2, 4, null, null, 4}), []*TreeNode{Ints2Tree([]int{4}), Ints2Tree([]int{2, 4})}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			res := findDuplicateSubtrees(ts.root)
			lenOut, lenIn := len(ts.output), len(res)
			ast.Equal(lenOut, lenIn)
			for j := 0; j < lenOut; j++ {
				ast.Equal(ts.output[j].PrintTreeNode(), res[j].PrintTreeNode())
			}
		})
	}
}
