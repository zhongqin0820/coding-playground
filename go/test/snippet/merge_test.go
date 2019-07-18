package snippet

import (
	"fmt"
	"testing"
)

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
	for that != nil && that.Next != nil {
		res += fmt.Sprintf("%d->", that.Val)
		that = that.Next
	}
	if that != nil {
		res += fmt.Sprintf("%d", that.Val)
	}
	return res
}
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}

func mergeLoop(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	res := &ListNode{Next: l1}
	ret := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node := l2.Next
			ret.Next, l2 = l2, l1.Next
			l2 = node
		} else {
			l1 = l1.Next
		}
		ret = ret.Next
	}
	if l1 != nil {
		ret.Next = l1.Next
	}
	if l2 != nil {
		ret.Next = l2.Next
	}
	return res.Next
}

func TestMerge(t *testing.T) {
	fmt.Printf("Hello World!\n")
	l1 := NewListNode([]int{1, 3, 5, 7, 9})
	l2 := NewListNode([]int{2, 4, 6, 8, 10})
	fmt.Println(l1.PrintList())
	fmt.Println(l2.PrintList())
	fmt.Println(merge(l1, l2).PrintList())
	fmt.Println(mergeLoop(l1, l2).PrintList())

}
