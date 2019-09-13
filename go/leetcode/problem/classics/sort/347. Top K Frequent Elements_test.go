package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	// 统计数频
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	n := len(m)
	// 对数频率排序
	counter := make([]int, 0, n)
	for _, v := range m {
		counter = append(counter, v)
	}
	// 从小到大
	sort.Ints(counter)
	// 数频下限
	min := counter[n-k]
	res := make([]int, 0, k)
	// 大于下限的都属于top k
	for key, v := range m {
		if v >= min {
			res = append(res, key)
		}
	}
	return res
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, topKFrequent(ts.nums, ts.k))
		})
	}
}
