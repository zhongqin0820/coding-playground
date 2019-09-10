package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-number-with-alternating-bits/description/
func hasAlternatingBits(n int) bool {
	// 计算低位的模板是0b01还是0b10
	std := n & 3
	if std != 1 && std != 2 {
		return false
	}
	// 用第一次得到模板去匹配后面的内容
	for n > 0 {
		if n&3 != std {
			return false
		}
		n >>= 2
	}
	return true
}

func hasAlternatingBitsII(n int) bool {
	a := (n ^ (n >> 1))
	return (a & (a + 1)) == 0
}

func TestHasAlternatingBits(t *testing.T) {
	tests := []struct {
		n      int
		output bool
	}{
		{5, true},
		{7, false},
		{11, false},
		{10, true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, hasAlternatingBits(ts.n))
			ast.Equal(ts.output, hasAlternatingBitsII(ts.n))
		})
	}
}
