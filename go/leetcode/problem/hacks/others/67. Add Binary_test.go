package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-binary/description/
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	// l是长度大的那个切片的长度
	l := len(a)
	// 将字符切片转换为整数切片
	isA := helper67trans(a, l)
	isB := helper67trans(b, l)
	// 直接利用整数切片进行进位操作，最终结果再转换回来
	return helper67makeString(helper67add(isA, isB))
}

func helper67trans(s string, l int) []int {
	res := make([]int, l)
	ls := len(s)
	for i, b := range s {
		res[l-ls+i] = int(b - '0')
	}
	return res
}

func helper67add(a, b []int) []int {
	// l+1，直接利用res[i]存放进位
	l := len(a) + 1
	res := make([]int, l)
	for i := l - 1; i >= 1; i-- {
		temp := res[i] + a[i-1] + b[i-1]
		res[i] = temp % 2
		res[i-1] = temp / 2
	}

	i := 0
	// i < l-1 而不是 < l 的原因是
	// "" + "" == "0"
	// 需要保留最后一个 '0'
	for i < l-1 && res[i] == 0 {
		i++
	}

	return res[i:]
}

// 将[]int转换为string
func helper67makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, addBinary(ts.a, ts.b))
		})
	}
}
