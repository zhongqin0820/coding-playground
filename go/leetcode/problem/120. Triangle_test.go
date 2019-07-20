package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/triangle/

// Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

// For example, given the following triangle
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

// Note:
// Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

// [i-1][j-1],[i-1][j]
//            [i][j]

func minimumTotal(triangle [][]int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n = len(triangle)
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			switch {
			case j == 0:
				triangle[i][j] += triangle[i-1][j]
			case j == i:
				triangle[i][j] += triangle[i-1][j-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	res := triangle[n-1][0]
	for j := 1; j < n; j++ {
		res = min(res, triangle[n-1][j])
	}
	return res
}

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minimumTotal(ts.input))
		})
	}
}
