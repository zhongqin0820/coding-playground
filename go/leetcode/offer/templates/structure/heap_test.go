package problem

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMinHeap(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, []int{1, 2, 3, 4, 5, 6}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			temp := intMinHeap(ts.input)
			h := &temp
			heap.Init(h)
			// 注意这里的初始长度指定为0，容量等于h.Len()
			res := make([]int, 0, h.Len())
			for h.Len() > 0 {
				res = append(res, heap.Pop(h).(int))
			}
			ast.Equal(ts.output, res)
		})
	}
}
