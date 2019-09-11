package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-consecutive-sequence/description/
// 应该使用并查集求解
func longestConsecutive(nums []int) int {
	var m map[int]int = make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			left, right := m[num-1], m[num+1]
			cur := left + 1 + right
			if cur > res {
				res = cur
			}
			m[num-left], m[num], m[num+right] = cur, cur, cur
		}
	}
	return res
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, longestConsecutive(ts.nums))
		})
	}
}
