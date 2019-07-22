package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/most-frequent-subtree-sum/

// Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	var pre func(*TreeNode) int
	pre = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			m[node.Val]++
			return node.Val
		} else if node.Left == nil {
			r := pre(node.Right)
			m[r+node.Val]++
			return r + node.Val
		} else if node.Right == nil {
			l := pre(node.Left)
			m[l+node.Val]++
			return l + node.Val
		}
		l := pre(node.Left)
		r := pre(node.Right)
		m[l+r+node.Val]++
		return l + r + node.Val
	}
	pre(root)

	max := -1
	res := []int{}
	for k, v := range m {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}

func TestFindFrequentTreeSum(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		output []int
	}{
		{Ints2Tree([]int{5, 2, -3}), []int{-3, 2, 4}},
		{Ints2Tree([]int{5, 2, -5}), []int{2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findFrequentTreeSum(ts.root))
		})
	}
}
