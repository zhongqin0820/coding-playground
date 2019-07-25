package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/integer-to-english-words/
// Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 231 - 1.

var lessThan21 = []string{
	"",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
	"Twenty",
}

var ten = []string{
	"",
	"",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var thousand = []string{
	"",
	"Thousand",
	"Million",
	"Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	i := 0

	for num > 0 {
		if num%1000 != 0 {
			res = helper273(num%1000) + thousand[i] + " " + res
		}

		num /= 1000
		i++
	}
	return strings.TrimRight(res, " ")
}

func helper273(num int) string {
	if num == 0 {
		return ""
	}

	if num <= 20 {
		return lessThan21[num] + " "
	}

	if num < 100 {
		return ten[num/10] + " " + helper273(num%10)
	}

	return lessThan21[num/100] + " Hundred " + helper273(num%100)
}

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		num    int
		output string
	}{
		{123, "One Hundred Twenty Three"},
		{50868, "Fifty Thousand Eight Hundred Sixty Eight"},
		{100, "One Hundred"},
		{30, "Thirty"},
		{21, "Twenty One"},
		{0, "Zero"},
		{123, "One Hundred Twenty Three"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{101, "One Hundred One"},
		{1, "One"},
		{11, "Eleven"},
		{111, "One Hundred Eleven"},
		{1111, "One Thousand One Hundred Eleven"},
		{11111, "Eleven Thousand One Hundred Eleven"},
		{111111, "One Hundred Eleven Thousand One Hundred Eleven"},
		{1111111, "One Million One Hundred Eleven Thousand One Hundred Eleven"},
		{11111111, "Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
		{111111111, "One Hundred Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
		{1111111111, "One Billion One Hundred Eleven Million One Hundred Eleven Thousand One Hundred Eleven"},
		{1000000000, "One Billion"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numberToWords(ts.num))
		})
	}
}
