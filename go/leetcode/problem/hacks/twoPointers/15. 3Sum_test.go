package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum/

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:
// The solution set must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			switch s := nums[l] + nums[i] + nums[r]; {
			case s < 0:
				l++
			case s > 0:
				r--
			default:
				res = append(res, []int{nums[l], nums[i], nums[r]})
				l, r = helper15(nums, l, r)
			}
		}
	}
	return res
}

func helper15(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{0, -1, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, threeSum(ts.input))
		})
	}
}
