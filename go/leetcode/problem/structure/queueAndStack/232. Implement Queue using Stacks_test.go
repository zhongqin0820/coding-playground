package problem

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/implement-queue-using-stacks/description/
type MyQueue struct {
	// s1负责入队列
	// s2辅助出队列，当栈s2为空时将栈s1中的元素倒到栈s2中再做出队列操作
	s1, s2 *list.List
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.PushBack(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if this.s2.Len() == 0 {
		for this.s1.Len() != 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	return this.s2.Remove(this.s2.Back()).(int)
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if this.s2.Len() == 0 {
		for this.s1.Len() != 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	return this.s2.Back().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.s1.Len() == 0 && this.s2.Len() == 0
}

func TestMyQueue(t *testing.T) {
	obj := ConstructorMyQueue()
	ast := assert.New(t)
	obj.Push(1)
	obj.Push(2)
	ast.Equal(1, obj.Peek())
	ast.Equal(1, obj.Peek())
	ast.Equal(false, obj.Empty())
}
