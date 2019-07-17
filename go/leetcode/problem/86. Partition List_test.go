package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/partition-list/

// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

// You should preserve the original relative order of the nodes in each of the two partitions.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	left, right := new(ListNode), new(ListNode)
	tempLeft, tempRight := left, right
	// 将链表拆成两部分，left部分是小于x，right部分大于x
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	// 将两部分结果拼接，注意需要将两部分的尾节点置nil
	left.Next, right.Next = nil, nil
	// 注意此处是判断右半部分是否不为空，不为空则拼接
	if tempRight.Next != nil {
		left.Next = tempRight.Next
	}
	return tempLeft.Next
}

// Example:

// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5
func TestPartition(t *testing.T) {
	tests := []struct {
		head   *ListNode
		x      int
		output *ListNode
	}{
		{nil, 0, nil},
		{NewListNode([]int{1, 4, 3, 2, 5, 2}), 3, NewListNode([]int{1, 2, 2, 4, 3, 5})},
		{NewListNode([]int{4, 3, 2, 1, 5, 2}), 3, NewListNode([]int{2, 1, 2, 4, 3, 5})},
		{NewListNode([]int{4, 3, 2, 1, 5, 2}), 0, NewListNode([]int{4, 3, 2, 1, 5, 2})},
		{NewListNode([]int{4, 3, 2, 1, 5, 2}), 10, NewListNode([]int{4, 3, 2, 1, 5, 2})},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			// ast.Equal(ts.output, partition(ts.head, ts.x))
			ast.Equal(ts.output.PrintList(), partition(ts.head, ts.x).PrintList())
		})
	}
}
