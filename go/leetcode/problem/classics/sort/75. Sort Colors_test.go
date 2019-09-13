package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-colors/description/
func sortColors(nums []int) {
	// 三路快排中的划分思路
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		switch nums[j] { // nums取值范围: [0,2]
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j+1
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			sortColors(ts.nums)
			ast.Equal(ts.output, ts.nums)
		})
	}
}
