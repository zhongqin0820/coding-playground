package containers

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

//我们自定义一个堆需要实现5个接口
//Len(),Less(),Swap()这是继承自sort.Interface
//Push()和Pop()是堆自已的接口

//返回长度
func (h *IntHeap) Len() int {
	return len(*h)
}

//比较大小(实现最小堆)
func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

//交换值
func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

//压入数据
func (h *IntHeap) Push(x interface{}) {
	//将数据追加到h中
	*h = append(*h, x.(int))
}

//弹出数据
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	//让h指向新的slice
	*h = old[0 : n-1]
	//返回最后一个元素
	return x
}

//打印堆
func (h *IntHeap) PrintHeap() {
	//元素的索引号
	i := 0
	//层级的元素个数
	levelCount := 1
	for i+1 <= h.Len() {
		fmt.Println((*h)[i : i+levelCount])
		i += levelCount
		if (i + levelCount*2) <= h.Len() {
			levelCount *= 2
		} else {
			levelCount = h.Len() - i
		}
	}
}

func TestHeap(t *testing.T) {
	a := IntHeap{6, 2, 3, 1, 5, 4}
	//初始化堆
	heap.Init(&a)
	a.PrintHeap()
	//弹出数据，保证每次操作都是规范的堆结构
	t.Log(heap.Pop(&a))
	a.PrintHeap()
	t.Log(heap.Pop(&a))
	a.PrintHeap()
	heap.Push(&a, 0)
	heap.Push(&a, 8)
	a.PrintHeap()
}
