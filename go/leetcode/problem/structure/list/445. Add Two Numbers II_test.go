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
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := make([]int, 0, 128), make([]int, 0, 128)
	for ; l1 != nil; l1 = l1.Next {
		s1 = append(s1, l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		s2 = append(s2, l2.Val)
	}
	sum, head := 0, &ListNode{Val: 0}
	for len(s1) != 0 || len(s2) != 0 {
		if len(s1) != 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		head.Val = sum % 10
		// ln是进位节点
		ln := &ListNode{Val: sum / 10}
		ln.Next = head
		head = ln
		// 此时的sum是进位值
		sum /= 10
	}

	if head.Val == 0 {
		return head.Next
	}
	return head
}

func TestAddTwoNumbersII(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{NewListNode([]int{7, 2, 4, 3}), NewListNode([]int{5, 6, 4}), NewListNode([]int{7, 8, 0, 7})},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, addTwoNumbersII(ts.l1, ts.l2))
		})
	}
}
