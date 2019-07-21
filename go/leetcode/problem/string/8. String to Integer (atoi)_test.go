package problem

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/string-to-integer-atoi/

// Implement atoi which converts a string to an integer.

// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

// If no valid conversion could be performed, a zero value is returned.

// Note:

// Only the space character ' ' is considered as whitespace character.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical value is out of the range of representable values, INT_MAX (2^31 − 1) or INT_MIN (−2^31) is returned.

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	var (
		sign int
		abs  string
	)
	// ignore the case of having prefix words
	switch {
	case strings.Contains("0123456789", str[:1]):
		sign, abs = 1, str
	case str[0] == '+':
		sign, abs = 1, str[1:]
	case str[0] == '-':
		sign, abs = -1, str[1:]
	}
	// skip the posfix words
	for i, _ := range abs {
		if abs[i] < '0' || '9' < abs[i] {
			abs = abs[:i]
			break
		}
	}
	// log.Println(abs)
	//
	res := 0
	for i, _ := range abs {
		res = res*10 + int(abs[i]-'0')
		switch {
		case sign == 1 && res > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && res*sign < math.MinInt32:
			return math.MinInt32

		}
	}
	return res * sign
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"", 0},
		{"-42", -42},
		{"   -42", -42},
		// Explanation: The first non-whitespace character is '-', which is the minus sign.
		// Then take as many numerical digits as possible, which gets 42.
		{"42", 42},
		{"     42", 42},
		{"4193 with words", 4193},
		// Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
		{"words and 987", 0},
		// Explanation: The first non-whitespace character is 'w', which is not a numerical
		// digit or a +/- sign. Therefore no valid conversion could be performed.
		{"-91283472332", -2147483648},
		// Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
		// Thefore INT_MIN (−2^31) is returned.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, myAtoi(ts.input))
		})
	}
}
