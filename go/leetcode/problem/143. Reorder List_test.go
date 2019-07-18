package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reorder-list/

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

// You may not modify the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reorderList(head *ListNode) {}
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//
	cur, Head, n := head, head, 0
	for cur != nil {
		cur = cur.Next
		n++
	}
	cur = head
	for i := 0; i < (n-1)/2; i++ {
		cur = cur.Next
	}
	// head -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
	//                   ^
	//                   |
	//                  cur

	// reverse cur...end
	pre, next := cur, cur.Next
	for next != nil {
		pre, next, next.Next = next, next.Next, pre
	}
	end := pre
	// head -> 1 -> 2 -> 3 <-> 4 <- 5 <- 6 <- end
	// merge
	for head != end {
		hNext, eNext := head.Next, end.Next
		head.Next, end.Next = end, hNext
		head, end = hNext, eNext
	}
	// in case of ring
	end.Next = nil
	return Head
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		head, output *ListNode
	}{
		{NewListNode([]int{1, 2, 3, 4}), NewListNode([]int{1, 4, 2, 3})},
		{NewListNode([]int{1, 2, 3, 4, 5}), NewListNode([]int{1, 5, 2, 4, 3})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output.PrintList(), reorderList(ts.head).PrintList())
		})
	}
}
