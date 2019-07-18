package snippet

import (
	"container/heap"
	"log"
	"testing"
)

func heapAdjust(a []int, i, n int) {
	var iChild int
	for ; 2*i+1 < n; i = iChild {
		iChild = 2*i + 1
		// get the max value out of i's child nodes
		if iChild+1 < n && a[iChild+1] > a[iChild] {
			iChild++
		}
		// get the max value between i and its child
		if a[i] < a[iChild] {
			a[i], a[iChild] = a[iChild], a[i]
		} else {
			break
		}
	}
}

func heapSort(a []int) {
	n := len(a)
	// build heap by using a
	for i := n / 2; i >= 0; i-- {
		heapAdjust(a, i, n)
	}
	// sort
	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapAdjust(a, 0, i)
	}
}

type MyHeap []int

func (mh *MyHeap) Len() int {
	return len(*mh)
}

func (mh *MyHeap) Less(i, j int) bool {
	return (*mh)[i] < (*mh)[j]
}

func (mh *MyHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MyHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MyHeap) Pop() interface{} {
	res := (*mh)[mh.Len()-1]
	(*mh) = (*mh)[:mh.Len()-1]
	return res
}

func TestHeapSort(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, -1}
	heapSort(v)
	log.Println(v)
	h := &MyHeap{1, 2, 3, 4, 5, 6, -1}
	res := make([]int, 0, h.Len())
	heap.Init(h)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	log.Println(res)
}

// go test -bench=. -run=^$
func BenchmarkHeapSort(b *testing.B) {
	b.ResetTimer()
	b.Run("heapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := []int{1, 2, 3, 4, 5, 6, -1}
			heapSort(v)
		}
	})

	b.ResetTimer()
	b.Run("heap.Pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := &MyHeap{1, 2, 3, 4, 5, 6, -1}
			res := make([]int, 0, h.Len())
			heap.Init(h)
			for h.Len() > 0 {
				res = append(res, heap.Pop(h).(int))
			}
		}
	})
	// BenchmarkHeapSort/heapSort-4            20000000            71.1 ns/op
	// BenchmarkHeapSort/heap.Pop-4             3000000           555 ns/op
}
