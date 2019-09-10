package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/missing-number/description/
func missingNumber(nums []int) int {
	// 这题还是相当于singleNumber
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	return res ^ len(nums)
}

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, missingNumber(ts.nums))
		})
	}
}
