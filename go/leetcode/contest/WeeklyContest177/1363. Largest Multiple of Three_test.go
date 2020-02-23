package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-multiple-of-three/
func largestMultipleOfThree(digits []int) string {
	sort.Ints(digits)
	if digits[len(digits)-1] == 0 {
		return "0"
	}
	a := make([]int, len(digits))
	one, two := 0, 0
	for i := range digits {
		a[i] = digits[i] % 3
		if a[i] == 1 {
			one++
		} else if a[i] == 2 {
			two++
		}
	}
	one, two = onetwo(one, two)
	res := ""
	for i := len(digits) - 1; i >= 0; i-- {
		if a[i] == 0 {
			res += string('0' + digits[i])
		} else if a[i] == 1 && one > 0 {
			res += string('0' + digits[i])
			one--
		} else if a[i] == 2 && two > 0 {
			res += string('0' + digits[i])
			two--
		}
	}
	return res
}

func onetwo(one, two int) (int, int) {
	len1, len2 := one%3, two%3
	if len1 == 0 && len2 == 2 {
		if one > 0 {
			len1 = 1
			len2 = 0
		}
	} else if len1 == 2 && len2 == 0 {
		if two > 0 {
			len1 = 0
			len2 = 0
		}
	} else if len1 > 0 && len2 > 0 {
		index := min(len1, len2)
		len1 -= index
		len2 -= index
	}

	return one - len1, two - len2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestLargestMultipleOfThree(t *testing.T) {
	tests := []struct {
		digits []int
		output string
	}{
		{[]int{8, 1, 9}, "981"},
		{[]int{8, 6, 7, 1, 0}, "8760"},
		{[]int{1}, ""},
		{[]int{0, 0, 0, 0, 0, 0}, "0"},
		{[]int{1, 1, 1}, "111"},
		{[]int{2, 2, 1, 1, 1}, "2211"},
		{[]int{7, 7, 7, 2, 2}, "7722"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, largestMultipleOfThree(ts.digits))
		})
	}
}
