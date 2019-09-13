package problem

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/min-stack/description/
type MinStack struct {
	// s1正常出入栈
	// s2存储最小值，每次入栈时与，若元素小于s2栈顶则压栈元素否则压栈s2栈顶元素
	s1, s2 *list.List
}

func ConstructorMinStack() MinStack {
	return MinStack{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.s1.PushBack(x)
	if this.s2.Len() == 0 {
		this.s2.PushBack(x)
	} else {
		min := this.GetMin()
		if x < min {
			this.s2.PushBack(x)
		} else {
			this.s2.PushBack(min)
		}
	}
}

func (this *MinStack) Pop() {
	this.s1.Remove(this.s1.Back())
	this.s2.Remove(this.s2.Back())
}

func (this *MinStack) Top() int {
	return this.s1.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.s2.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func TestMinStack(t *testing.T) {
	ast := assert.New(t)
	obj := ConstructorMinStack()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	ast.Equal(-3, obj.GetMin())
	obj.Pop()
	ast.Equal(0, obj.Top())
	ast.Equal(-2, obj.GetMin())
}
