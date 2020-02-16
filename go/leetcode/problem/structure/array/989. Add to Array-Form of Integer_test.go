package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-to-array-form-of-integer/
func addToArrayForm(A []int, K int) []int {
	B := make([]int, 0, 10)
	for ; K != 0; K /= 10 {
		B = append(B, K%10)
	}
	for i, j := 0, len(B)-1; i < j; i, j = i+1, j-1 {
		B[i], B[j] = B[j], B[i]
	}
	a, b, carry := len(A)-1, len(B)-1, 0
	res := make([]int, 0, 1024)
	for ; a >= 0 && b >= 0; a, b = a-1, b-1 {
		temp := A[a] + B[b] + carry
		carry = temp / 10
		res = append(res, temp%10)
	}
	for ; a >= 0; a-- {
		temp := A[a] + carry
		carry = temp / 10
		res = append(res, temp%10)
	}
	for ; b >= 0; b-- {
		temp := B[b] + carry
		carry = temp / 10
		res = append(res, temp%10)
	}
	if carry != 0 {
		res = append(res, carry)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func TestAddToArrayForm(t *testing.T) {
	tests := []struct {
		A      []int
		K      int
		output []int
	}{
		{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0}, 806, []int{8, 0, 6}},
		{[]int{0}, 23, []int{2, 3}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, addToArrayForm(ts.A, ts.K))
		})
	}
}
