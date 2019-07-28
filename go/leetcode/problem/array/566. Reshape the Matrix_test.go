package problem

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reshape-the-matrix/description/

// In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.
// You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.
// The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.
// If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

// Note:
// The height and width of the given matrix is in range [1, 100].
// The given r and c are all positive.
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil || len(nums) == 0 || len(nums[0]) == 0 {
		return nums
	}
	n, m := len(nums), len(nums[0])
	if (n <= r && m <= c) || (n*m < r*c) {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	//
	iRes, jRes := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			log.Println(nums[i][j])
			res[iRes][jRes] = nums[i][j]
			jRes++
			if jRes == c {
				jRes = 0
				iRes++
			}
		}
	}
	return res
}

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		nums   [][]int
		r, c   int
		output [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		// Explanation:
		// The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
		// Explanation:
		// There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, matrixReshape(ts.nums, ts.r, ts.c))
		})
	}
}
