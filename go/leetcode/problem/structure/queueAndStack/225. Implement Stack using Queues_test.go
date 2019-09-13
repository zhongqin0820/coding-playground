package problem

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/implement-stack-using-queues/description/
type MyStack struct {
	// q1负责入栈
	// q2辅助出栈，每次出栈都将q1的元素倒到q2中直到剩下最后一个元素，将其出栈后再将q2中元素倒回q1中
	q1, q2 *list.List
}

func ConstructorMyStack() MyStack {
	return MyStack{
		q1: list.New(),
		q2: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.q1.PushBack(x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	for this.q1.Len() > 1 {
		this.q2.PushBack(this.q1.Remove(this.q1.Front()))
	}
	e := this.q1.Remove(this.q1.Front()).(int)
	for this.q2.Len() > 0 {
		this.q1.PushBack(this.q2.Remove(this.q2.Front()))
	}
	return e
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	for this.q1.Len() > 1 {
		this.q2.PushBack(this.q1.Remove(this.q1.Front()))
	}
	e := this.q1.Front().Value.(int)
	this.q2.PushBack(this.q1.Remove(this.q1.Front()))
	for this.q2.Len() > 0 {
		this.q1.PushBack(this.q2.Remove(this.q2.Front()))
	}
	return e
}

func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0 && this.q2.Len() == 0
}

func TestMyStack(t *testing.T) {
	ast := assert.New(t)
	obj := ConstructorMyStack()
	obj.Push(1)
	obj.Push(2)
	ast.Equal(2, obj.Top())
	ast.Equal(2, obj.Pop())
	ast.Equal(false, obj.Empty())
}
