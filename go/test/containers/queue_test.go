package containers

import (
	"container/list"
	"testing"
)

// description: https://www.cnblogs.com/zmlctt/p/3985128.html
type MyDeQueue struct {
	s1 *list.List
	s2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyDeQueue {
	return MyDeQueue{s1: list.New(), s2: list.New()}
}

/** Push element x to the front of queue. */
func (this *MyDeQueue) PushFront(x int) {
	this.s2.PushBack(x)
}

/** Push element x to the back of queue. */
func (this *MyDeQueue) PushBack(x int) {
	this.s1.PushBack(x)
}

/** Get the front element. */
func (this *MyDeQueue) Front() int {
	if this.Empty() {
		return -1
	}
	if this.s2.Len() == 0 {
		return this.Back()
	}
	res := this.s2.Back().Value.(int)
	return res
}

/** Get the back element. */
func (this *MyDeQueue) Back() int {
	if this.Empty() {
		return -1
	}
	if this.s1.Len() == 0 {
		return this.Front()
	}
	res := this.s1.Back().Value.(int)
	return res
}

/** Removes the element from the front of queue and returns that element. */
func (this *MyDeQueue) PopFront() int {
	if this.Empty() {
		return -1
	}
	if this.s2.Len() == 0 {
		return this.PopBack()
	}
	res := this.s2.Back().Value.(int)
	this.s2.Remove(this.s2.Back())
	return res
}

/** Removes the element from the back of queue and returns that element. */
func (this *MyDeQueue) PopBack() int {
	if this.Empty() {
		return -1
	}
	if this.s1.Len() == 0 {
		return this.PopFront()
	}
	res := this.s1.Back().Value.(int)
	this.s1.Remove(this.s1.Back())
	return res
}

/** Returns whether the queue is empty. */
func (this *MyDeQueue) Empty() bool {
	return this.s1.Len() == 0 && this.s2.Len() == 0
}

func TestMyDequeue(t *testing.T) {
	obj := Constructor()
	obj.PushBack(1)
	obj.PushBack(2)
	obj.PushFront(3)
	obj.PushFront(4)
	if obj.Front() != 4 {
		t.Fatal("the front is: ", obj.Front())
	}
	t.Log(obj.Back())
	t.Log(obj.PopBack())
	t.Log(obj.Back())
}
