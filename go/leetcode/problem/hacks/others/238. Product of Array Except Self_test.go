package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/product-of-array-except-self/description/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	// 左侧元素乘积
	left := 1
	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}
	// 右侧元素乘积
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, productExceptSelf(ts.nums))
		})
	}
}
