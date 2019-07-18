package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

// Given a sorted linked list, delete all duplicates such that each element appear only once.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := new(ListNode)
	ret, cur := head, head.Next
	res.Next = ret
	for cur != nil {
		if ret.Val == cur.Val {
			ret.Next = cur.Next
			cur = cur.Next
			continue
		}
		cur = cur.Next
		ret = ret.Next
	}
	return res.Next
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input, output *ListNode
	}{
		{NewListNode([]int{1, 1, 2}), NewListNode([]int{1, 2})},
		{NewListNode([]int{1, 1, 2, 3, 3}), NewListNode([]int{1, 2, 3})},
		{nil, nil},
		{NewListNode([]int{}), NewListNode([]int{})},
		{NewListNode([]int{1}), NewListNode([]int{1})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), deleteDuplicates(ts.input).PrintList())
		})
	}
}
