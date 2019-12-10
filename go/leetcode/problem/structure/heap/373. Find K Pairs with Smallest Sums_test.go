package problem

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct {
	i, j int
	sum  int
}

type pq []*pair

func (h pq) Len() int           { return len(h) }
func (h pq) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h pq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pq) Push(x interface{}) {
	e := x.(*pair)
	*h = append(*h, e)
}
func (h *pq) Pop() interface{} {
	n := len(*h)
	e := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return e
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 边界检测
	res := make([][]int, 0, k)
	if nums1 == nil || len(nums1) == 0 ||
		nums2 == nil || len(nums2) == 0 {
		return res
	}
	// 初始化
	q := make(pq, helper373Min(k, len(nums1)))
	for i := 0; i < k && i < len(nums1); i++ {
		q[i] = &pair{
			i:   i,
			j:   0,
			sum: nums1[i] + nums2[0],
		}
	}
	heap.Init(&q)
	// 遍历压入最小到结果
	var min *pair
	for ; k > 0 && len(q) > 0; k-- {
		min = heap.Pop(&q).(*pair)
		res = append(res, []int{nums1[min.i], nums2[min.j]})
		// 说明min是当前最小，那么，min.j+1相较其它队列中的下一个元素更符合
		if min.j+1 < len(nums2) {
			heap.Push(&q, &pair{
				i:   min.i,
				j:   min.j + 1,
				sum: nums1[min.i] + nums2[min.j+1],
			})
		}
	}
	return res
}

func helper373Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		k      int
		output [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{[]int{1, 2}, []int{1, 4}, []int{1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{[]int{1, 1}, []int{1, 1}}},
		{[]int{1, 2}, []int{3}, 3, [][]int{[]int{1, 3}, []int{2, 3}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 10, [][]int{[]int{1, 1}, []int{1, 1}, []int{1, 2}, []int{1, 2}, []int{2, 1}, []int{1, 3}, []int{2, 2}, []int{1, 3}, []int{2, 3}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, kSmallestPairs(ts.nums1, ts.nums2, ts.k))
		})
	}
}
