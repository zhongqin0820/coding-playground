package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/single-number-iii/description/
func singleNumberIII(nums []int) []int {
	// xor = a ^ b
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	// 得到a与b最右侧不相同的位
	xor &= -xor
	// 迭代区分出对应的数
	var a, b int
	for _, num := range nums {
		if num&xor == 0 { //因为相同数的个数是2个，所以此处可以按照这种方式分
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

func TestSingleNumberIII(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{5, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, singleNumberIII(ts.nums))
		})
	}
}
