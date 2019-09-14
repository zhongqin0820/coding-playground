package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-strings/description/
func addStrings(num1 string, num2 string) string {
	// 找到最长的那个字符串的长度
	l, b := len(num1), len(num2)
	if l < b {
		l = b
	}
	// 将字符串转换为切片
	nums1, nums2 := helper415Trans(num1, l), helper415Trans(num2, l)
	// 直接在[]int上做add
	return helper415String(helper415Add(nums1, nums2))
}

func helper415Trans(s string, l int) []int {
	var a []int = make([]int, l)
	ls := len(s)
	for i := range s {
		a[l-ls+i] = int(s[i] - '0')
	}
	return a
}

func helper415Add(a, b []int) []int {
	l := len(a) + 1
	var res []int = make([]int, l)

	for i := l - 1; i >= 1; i-- {
		temp := a[i-1] + b[i-1] + res[i]
		res[i] = temp % 10   //当前值
		res[i-1] = temp / 10 //进位
	}
	// 忽略开头的0
	var i int
	for ; i < l-1 && res[i] == 0; i++ {
	}
	return res[i:]
}

func helper415String(a []int) string {
	var bs = make([]byte, len(a))
	for i := range bs {
		bs[i] = byte(a[i]) + '0'
	}
	return string(bs)
}

func TestAddStrings(t *testing.T) {
	tests := []struct {
		num1   string
		num2   string
		output string
	}{
		{"123", "1", "124"},
		{"0", "0", "0"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, addStrings(ts.num1, ts.num2))
		})
	}
}
