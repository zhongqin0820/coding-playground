package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/non-overlapping-intervals/description/
type intervalSlice []struct {
	start, end int
}

func (i intervalSlice) Len() int {
	return len(i)
}

func (i intervalSlice) Less(j, k int) bool {
	return i[j].end < i[k].end
}

func (i intervalSlice) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}

func eraseOverlapIntervals(intervals [][]int) int {
	end, res, n := -1<<31, 0, len(intervals)
	interval := make(intervalSlice, n)
	for i := 0; i < n; i++ {
		interval[i].start, interval[i].end = intervals[i][0], intervals[i][1]
	}
	// interval根据结束值升序排序
	sort.Sort(interval)
	// 先安排结束值小的区间
	for _, inter := range interval {
		if end <= inter.start {
			end = inter.end
		} else {
			res++
		}
	}
	return res
}

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		output    int
	}{
		{
			[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}},
			1,
		},
		{
			[][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}},
			2,
		},
		{
			[][]int{[]int{1, 2}, []int{2, 3}},
			0,
		},
		{
			[][]int{[]int{1, 2}},
			0,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, eraseOverlapIntervals(ts.intervals))
		})
	}
}
