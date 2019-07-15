package problem

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(a []int) *ListNode {
	if a == nil || len(a) == 0 {
		return nil
	}
	//
	node := new(ListNode)
	node.Val = a[0]
	head := node
	for i := 1; i < len(a); i++ {
		temp := &ListNode{
			Val:  a[i],
			Next: nil,
		}
		node.Next = temp
		node = node.Next
	}
	return head
}

func (this *ListNode) PrintList() string {
	that := this
	res := ""
	for that.Next != nil {
		res += fmt.Sprintf("%d->", that.Val)
		that = that.Next
	}
	if that != nil {
		res += fmt.Sprintf("%d", that.Val)
	}
	return res
}

func (this *ListNode) Compare(node *ListNode) bool {
	that := this
	for node != nil && that != nil {
		if node.Val != that.Val {
			return false
		}
		node = node.Next
		that = that.Next
	}
	return node == nil && that == nil
}
