package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-sorted-lists/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	ret := head
	// 链不为空
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			ret.Next = l1
			l1 = l1.Next
		} else {
			ret.Next = l2
			l2 = l2.Next
		}
		ret = ret.Next
	}
	// 剩下的内容
	if l1 == nil {
		ret.Next = l2
	}
	if l2 == nil {
		ret.Next = l1
	}
	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{NewListNode([]int{1, 2, 4}), NewListNode([]int{1, 3, 4}), NewListNode([]int{1, 1, 2, 3, 4, 4})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, mergeTwoLists(ts.l1, ts.l2))
		})
	}
}
