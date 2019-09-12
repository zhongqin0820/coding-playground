package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/set-mismatch/description/

// The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.
// Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

// Note:
// The given array size will in the range [2, 10000].
// The given array's numbers won't have any order.

func findErrorNums(nums []int) []int {
	if nums == nil || len(nums) == 1 {
		return nums
	}
	dup := 0
	for i := 0; i < len(nums); i++ {
		idx := helper645(nums[i]) // nums[i] 属于[1,len(nums)]

		if nums[idx-1] < 0 {
			dup = idx
		} else {
			nums[idx-1] = -nums[idx-1]
		}
	}

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}
	return []int{dup, mis}
}

func helper645(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findErrorNums(ts.nums))
		})
	}
}
