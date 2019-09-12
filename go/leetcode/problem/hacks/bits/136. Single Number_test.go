package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/single-number/description/
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, singleNumber(ts.nums))
		})
	}
}
