package problem

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	i, j := 0, col-1
	for i < row && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

// 算法改进
func searchMatrixAdv(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	for i := 0; i < row; i++ {
		// 找到目标行，对目标行使用二分搜索（binary searchs）
		if matrix[i][col-1] >= target {
			return bs(matrix[i], target, 0, col)
		}
	}
	return false
}

// bs 二分搜索（binary searchs）
func bs(array []int, target, low int, high int) bool {
	if low > high {
		return false
	}
	mid := (low + high) / 2
	if array[mid] == target {
		return true
	}
	if array[mid] < target {
		return bs(array, target, mid+1, high)
	}
	return bs(array, target, low, mid-1)
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false},
		{[][]int{}, 0, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d Naive", i+1), func(t *testing.T) {
			if res := searchMatrix(ts.matrix, ts.target); res != ts.output {
				t.Errorf("expected %t got %t", ts.output, res)
			}
		})
		t.Run(fmt.Sprintf("Example %d Adv", i+1), func(t *testing.T) {
			if res := searchMatrixAdv(ts.matrix, ts.target); res != ts.output {
				t.Errorf("expected %t got %t", ts.output, res)
			}
		})
	}
}
