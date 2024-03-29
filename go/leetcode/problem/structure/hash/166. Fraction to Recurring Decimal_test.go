package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/fraction-to-recurring-decimal/

// Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

// If the fractional part is repeating, enclose the repeating part in parentheses.

func fractionToDecimal(n int, d int) string {
	if n == 0 {
		return "0"
	}

	if n*d < 0 {
		// n, d 异号，则结果前面有负号
		return "-" + fractionToDecimal(helper166(n), helper166(d))
	}

	// 确保 n  和 d 是非负数
	n, d = helper166(n), helper166(d)

	if n >= d {
		// n / d 的小数部分
		ds := fractionToDecimal(n%d, d)
		return strconv.Itoa(n/d) + ds[1:]
	}

	// digits 用来保存 n/d 的结果
	digits := make([]byte, 2, 1024)
	digits[0] = '0'
	digits[1] = '.'
	// idx 是 n/d 的结果的在 digits 的索引号
	idx := 2
	// rec[n] = idx
	rec := make(map[int]int, 1024)
	for {
		if i, ok := rec[n]; ok {
			// n 重复出现，则说明 n/d 是下一个循环的开始
			// 循环部分的起点，就是 n 上次出现的 idx 值
			return fmt.Sprintf("%s(%s)", string(digits[:i]), string(digits[i:]))
		}

		rec[n] = idx

		n *= 10
		idx++

		digits = append(digits, byte(n/d)+'0')
		n %= d

		if n == 0 {
			// 不会有循环部分
			return string(digits)
		}
	}
}

func helper166(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func TestFractionToDecimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
		output      string
	}{
		{1, 2, "0.5"},
		{2, 1, "2"},
		{2, 3, "0.(6)"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, fractionToDecimal(ts.numerator, ts.denominator))
		})
	}
}
