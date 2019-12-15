package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-167/problems/convert-binary-number-in-a-linked-list-to-integer/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	var dfs func(*ListNode)
	var res, exp = 0, 1
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		res += node.Val * exp
		exp *= 2
	}
	dfs(head)
	return res
}

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		head   *ListNode
		output int
	}{
		{NewListNode([]int{1, 0, 1}), 5},
		{NewListNode([]int{0}), 0},
		{NewListNode([]int{1}), 1},
		{NewListNode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), 18880},
		{NewListNode([]int{0, 0}), 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getDecimalValue(ts.head))
		})
	}
}
