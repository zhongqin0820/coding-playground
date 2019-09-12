package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-number/

// Validate if a given string can be interpreted as a decimal number.
// Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one. However, here is a list of characters that can be in a valid decimal number:

// Numbers 0-9
// Exponent - "e"
// Positive/negative sign - "+"/"-"
// Decimal point - "."
// Of course, the context of these characters also matters in the input.

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	numbers := strings.Split(s, "e")
	if len(numbers) == 1 {
		// 非科学计数的数
		return helper65IsFloat(numbers[0])
	}
	if len(numbers) == 2 {
		// 科学计数的数，前一半可以是浮点数，后一半必须是整数
		return helper65IsFloat(numbers[0]) && helper65IsInt(numbers[1])
	}
	return false
}

func helper65IsFloat(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 有符号
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	numbers := strings.Split(s, ".")
	// 不是浮点数，则判断是否是整数
	if len(numbers) == 1 {
		return helper65IsUInt(numbers[0])
	}
	// 不合法的数，例如：0.2.3
	if len(numbers) > 2 {
		return false
	}
	// 是浮点数，则“.”前后都得是整数
	// 例如：.3
	if len(numbers[0]) == 0 {
		return helper65IsUInt(numbers[1])
	}
	// 例如：2.
	if len(numbers[1]) == 0 {
		return helper65IsUInt(numbers[0])
	}
	// 例如：2.3
	return helper65IsUInt(numbers[0]) && helper65IsUInt(numbers[1])
}

func helper65IsInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 有符号
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	// 剩下的是无符号部分
	return helper65IsUInt(s)
}

func helper65IsUInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 判断是否有非法字符，存在非法字符返回错误
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{
		{"0", true},
		{" 0.1", true},
		{"abc", false},
		{"1  a", false},
		{"2e10", true},
		{"  -90e3", true},
		{" 1e", false},
		{"e3", false},
		{"6e-1", true},
		{"99e2.5", false},
		{"53.5e93", true},
		{" --6", false},
		{"-+3", false},
		{"95a54e53", false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isNumber(ts.s))
		})
	}
}
