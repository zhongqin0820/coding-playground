package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/

// Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.
// Note that it is the kth smallest element in the sorted order, not the kth distinct element.

// Note:
// You may assume k is always valid, 1 ≤ k ≤ n^2.

func kthSmallest(matrix [][]int, k int) int {
	if matrix == nil || len(matrix) == 0 {
		return -1
	}
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]

	for low < high {
		mid := low + (high-low)/2
		count := 0
		j := n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += j + 1
		}

		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		matrix [][]int
		k      int
		output int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, kthSmallest(ts.matrix, ts.k))
		})
	}
}
