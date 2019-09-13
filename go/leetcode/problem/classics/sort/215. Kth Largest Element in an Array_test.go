package problem

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
// 使用堆结构，小顶堆找top k大，结果在堆顶
func findKthLargest(nums []int, k int) int {
	h := &intMinHeap{}
	heap.Init(h)
	for _, num := range nums {
		// 插入数据
		heap.Push(h, num)
		// 弹出数据，调整堆
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

type intMinHeap []int

func (hs intMinHeap) Len() int { return len(hs) }

func (hs intMinHeap) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }

func (hs intMinHeap) Less(i, j int) bool { return hs[i] < hs[j] }

func (hs *intMinHeap) Push(e interface{}) { *hs = append(*hs, e.(int)) }

func (hs *intMinHeap) Pop() interface{} {
	e := (*hs)[hs.Len()-1]
	*hs = (*hs)[:hs.Len()-1]
	return e
}

// 使用快速排序的partition
func findKthLargestII(nums []int, k int) int {
	l, h, k := 0, len(nums)-1, len(nums)-k
	for l < h {
		// partition
		j := helper215(nums, l, h)
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			h = j - 1
		}
	}
	return nums[k]
}

func helper215(nums []int, l, h int) int {
	i, j := l, h+1
	for {
		// pivot key == nums[l]
		for i++; nums[i] < nums[l] && i < h; i++ {
		}
		for j--; nums[j] > nums[l] && j > l; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findKthLargest(ts.nums, ts.k))
			ast.Equal(ts.output, findKthLargestII(ts.nums, ts.k))
		})
	}
}
