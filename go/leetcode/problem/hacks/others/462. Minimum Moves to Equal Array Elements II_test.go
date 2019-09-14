package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/description/
func minMoves2(nums []int) int {
	sort.Ints(nums)
	res, l, h := 0, 0, len(nums)-1
	for l <= h {
		res += nums[h] - nums[l]
		l, h = l+1, h-1
	}
	return res
}

func TestMinMoves2(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 2, 3}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minMoves2(ts.nums))
		})
	}
}
