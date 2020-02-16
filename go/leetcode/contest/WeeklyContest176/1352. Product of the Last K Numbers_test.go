package contest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/product-of-the-last-k-numbers/
type ProductOfNumbers struct {
	nums  []int
	zeros int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		nums:  []int{1},
		zeros: 0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if n := len(this.nums) - 1; num != 0 {
		this.nums = append(this.nums, this.nums[n]*num)
	} else {
		this.zeros = n + 1
		this.nums = append(this.nums, 1)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if n := len(this.nums); this.zeros > 0 && n-this.zeros <= k {
		return 0
	} else {
		return this.nums[n-1] / this.nums[n-k-1]
	}
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
func TestProductOfNumbers(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	ast := assert.New(t)
	// 3,0,2,5,4
	ast.Equal(20, obj.GetProduct(2))
	ast.Equal(40, obj.GetProduct(3))
	ast.Equal(0, obj.GetProduct(4))
	obj.Add(8)
	// 3,0,2,5,4,8
	ast.Equal(32, obj.GetProduct(2))
}
