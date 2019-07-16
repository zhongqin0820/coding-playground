package problem

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// new a link list via a 1D slice
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

// print the link list for test usage
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

// compare two linked list
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

// compare two 2D slice
func slice2DCompare(a, b *[][]int) bool {
	if a == nil || b == nil || *a == nil || *b == nil {
		return false
	}
	if len((*a)) == 0 && len((*a)) == len((*b)) {
		return true
	}
	if len((*a)) != len((*b)) || len((*a)[0]) != len((*b)[0]) {
		return false
	}
	for i, v := range *a {
		for j, val := range (*b)[i] {
			if v[j] != val {
				return false
			}
		}
	}
	return true
}

// compare two 1D slice
func slice1DCompare(a, b []int) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(a) == len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
