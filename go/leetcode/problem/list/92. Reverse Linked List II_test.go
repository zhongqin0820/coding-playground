package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-linked-list-ii/

// Reverse a linked list from position m to n. Do it in one-pass.

// Note: 1 ≤ m ≤ n ≤ length of list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	res := new(ListNode)
	res.Next = head
	pre, cur := res, res

	// find the head of the sub-list
	for i := 0; i < m; i++ {
		pre = cur
		cur = cur.Next
	}
	tail := cur
	cur = cur.Next
	// pre -> head ... tail -> cur -> curNext
	//          |  ...         |
	//          m  ...         i
	for i := m; i < n; i++ {
		head = pre.Next
		pre.Next = cur

		curNext := cur.Next
		cur.Next = head

		cur = curNext
		tail.Next = curNext
	}
	return res.Next
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		input  *ListNode
		m      int
		n      int
		output *ListNode
	}{
		{nil, 0, 0, nil},
		{NewListNode([]int{1}), 0, 1, NewListNode([]int{1})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 5, 5, NewListNode([]int{1, 2, 3, 4, 5})},
		{NewListNode([]int{1, 2, 3, 4, 5}), 2, 4, NewListNode([]int{1, 4, 3, 2, 5})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			// ast.Equal(ts.output, reverseBetween(ts.input, ts.m, ts.n))
			ast.Equal(ts.output.PrintList(), reverseBetween(ts.input, ts.m, ts.n).PrintList())
		})
	}
}
