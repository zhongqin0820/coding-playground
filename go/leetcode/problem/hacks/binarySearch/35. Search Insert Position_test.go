package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i, j, m := 0, len(nums)-1, 0
	for i <= j {
		m = i + (j-i)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{nil, 0, 0},
		{[]int{}, 0, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, searchInsert(ts.nums, ts.target))
		})
	}
}
