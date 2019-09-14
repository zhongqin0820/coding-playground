package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/majority-element/description/
// Boyer-Moore Majority Vote Algorithm
func majorityElement(nums []int) int {
	res, t := nums[0], 1
	for i := 1; i < len(nums); i++ {
		switch {
		case res == nums[i]:
			t++
		case t > 0:
			t--
		default:
			res = nums[i]
			t = 1
		}
	}
	return res
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, majorityElement(ts.nums))
		})
	}
}
