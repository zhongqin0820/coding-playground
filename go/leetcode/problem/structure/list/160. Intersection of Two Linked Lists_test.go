package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	temp1, temp2 := headA, headB
	for temp1 != temp2 {
		// a+c+b
		if temp1 == nil {
			temp1 = headB
		} else {
			temp1 = temp1.Next
		}
		// a+b+c
		if temp2 == nil {
			temp2 = headA
		} else {
			temp2 = temp2.Next
		}
	}
	return temp1
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		headA, headB *ListNode
		output       *ListNode
	}{
		// 这道题的样例很奇怪，已经AC
		{NewListNode([]int{4, 1, 8, 4, 5}), NewListNode([]int{5, 0, 1, 8, 4, 5}), NewListNode([]int{})}, //8
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, getIntersectionNode(ts.headA, ts.headB))
		})
	}
}
