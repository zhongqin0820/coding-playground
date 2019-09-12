package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/binary-trees-with-factors/

// Given an array of unique integers, each integer is strictly greater than 1.
// We make a binary tree using these integers and each number may be used for any number of times.
// Each non-leaf node's value should be equal to the product of the values of it's children.
// How many binary trees can we make?  Return the answer modulo 10 ** 9 + 7.

// Note:
// 1 <= A.length <= 1000.
// 2 <= A[i] <= 10 ^ 9.

func numFactoredBinaryTrees(A []int) int {
	sort.Ints(A)
	factorIndex := make(map[int]int, len(A))
	for i := range A {
		factorIndex[A[i]] = i
	}

	ress := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		ress[i] = 1
		for j := 0; j < i; j++ {
			quotient, remainder := A[i]/A[j], A[i]%A[j]
			k, isFactor := factorIndex[quotient]

			if remainder != 0 || !isFactor {
				continue
			}
			ress[i] += ress[j] * ress[k]
		}
		ress[i] = ress[i] % (1E9 + 7)
	}

	res := 0
	for _, v := range ress {
		res += v
	}
	return res % (1E9 + 7)
}

func TestNumFactoredBinaryTrees(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{2, 4}, 3},
		{[]int{2, 4, 5, 10}, 7},
		{[]int{46, 144, 5040, 4488, 544, 380, 4410, 34, 11, 5, 3063808, 5550, 34496, 12, 540, 28, 18, 13, 2, 1056, 32710656, 31, 91872, 23, 26, 240, 18720, 33, 49, 4, 38, 37, 1457, 3, 799, 557568, 32, 1400, 47, 10, 20774, 1296, 9, 21, 92928, 8704, 29, 2162, 22, 1883700, 49588, 1078, 36, 44, 352, 546, 19, 523370496, 476, 24, 6000, 42, 30, 8, 16262400, 61600, 41, 24150, 1968, 7056, 7, 35, 16, 87, 20, 2730, 11616, 10912, 690, 150, 25, 6, 14, 1689120, 43, 3128, 27, 197472, 45, 15, 585, 21645, 39, 40, 2205, 17, 48, 136}, 509730797},
		{[]int{18, 3, 6, 2}, 12},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numFactoredBinaryTrees(ts.input))
		})
	}
}
