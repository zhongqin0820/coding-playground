package problem

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/gray-code/

// The gray code is a binary numeral system where two successive values differ in only one bit.

// Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0, 1}
	base := 2
	for i := 1; i < n; i++ {
		rlen := len(res)
		// 前半部分一样，只需添加后半部分
		// 而后半部分的计算是基于前半部分+高位为1（即2^i）
		for j := rlen - 1; j >= 0; j-- {
			res = append(res, res[j]+base)
		}
		// 更新base = 2^(i+1)
		base = base * 2
	}
	return res
}

func TestGrayCode(t *testing.T) {
	tests := []struct {
		input  int
		output []int
	}{
		{2, []int{0, 1, 3, 2}},
		// Explanation:
		// 00 - 0
		// 01 - 1
		// 11 - 3
		// 10 - 2

		// For a given n, a gray code sequence may not be uniquely defined.
		// For example, [0,2,3,1] is also a valid gray code sequence.

		// 00 - 0
		// 10 - 2
		// 11 - 3
		// 01 - 1
		{0, []int{0}},
		// Explanation: We define the gray code sequence to begin with 0.
		//              A gray code sequence of n has size = 2n, which for n = 0 the size is 20 = 1.
		//              Therefore, for n = 0 the gray code sequence is [0].
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			if res := grayCode(ts.input); !slice1DCompare(res, ts.output) {
				t.Errorf("expected %v\n", ts.output)
				t.Errorf("got %v\n", res)
			}
		})
	}
}
