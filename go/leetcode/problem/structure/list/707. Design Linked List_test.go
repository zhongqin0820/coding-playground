package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/design-linked-list/
type MyLinkedList struct {
	size       int
	head, tail *node
}

type node struct {
	val  int
	next *node
}

// Constructor initialize your data structure here. */
func Constructor() MyLinkedList {
	t := &node{}
	h := &node{next: t}
	// 把 head 和 tail 分别用空 node 单独表示
	// 会让 AddAtHead 和 AddAtTail 的逻辑非常简单
	// 可以试着让 size = 0 的时候， head 与 tail 为 nil
	// 看看程序的结构有多复杂
	return MyLinkedList{
		head: h,
		tail: t,
	}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || l.size <= index {
		return -1
	}
	i, cur := 0, l.head.next
	for i < index {
		i++
		cur = cur.next
	}
	return cur.val
}

// AddAtHead a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (l *MyLinkedList) AddAtHead(val int) {
	nd := &node{val: val}
	nd.next = l.head.next
	l.head.next = nd
	l.size++
}

// AddAtTail append a node of value val to the last element of the linked list.
func (l *MyLinkedList) AddAtTail(val int) {
	l.tail.val = val
	nd := &node{}
	l.tail.next = nd
	l.tail = nd
	l.size++
}

// AddAtIndex add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index < 0 || l.size < index:
		return
	case index == 0:
		l.AddAtHead(val)
		return
	case index == l.size:
		l.AddAtTail(val)
		return
	}

	i, cur := -1, l.head
	for i+1 < index {
		i++
		cur = cur.next
	}

	nd := &node{val: val}
	nd.next = cur.next
	cur.next = nd

	l.size++
}

// DeleteAtIndex delete the index-th node in the linked list, if the index is valid.
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || l.size <= index {
		return
	}

	i, cur := -1, l.head
	for i+1 < index {
		i++
		cur = cur.next
	}

	cur.next = cur.next.next

	l.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func TestMyLinkedList(t *testing.T) {
	ast := assert.New(t)
	obj := Constructor()
	// 1 -> 3
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	ast.Equal(1, obj.Get(0))
	ast.Equal(3, obj.Get(1))
	// 1 -> 2 -> 3
	obj.AddAtIndex(1, 2)
	ast.Equal(2, obj.Get(1))
	obj.DeleteAtIndex(1)
	// 1 -> 3
	ast.Equal(3, obj.Get(1))
}
