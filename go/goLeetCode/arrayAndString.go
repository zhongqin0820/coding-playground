package main

import (
	"fmt"
	"log"
)

//[[1,2,3],[4,5,6],[7,8,9]]
func exp2() {
	// res := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	res := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	log.Printf("%v\n", res)
}

// right down left up
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var Res []int
	var left, right, up, down = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	for c, dir := 0, 0; left <= right && up <= down; c++ {
		switch dir = c % 4; dir {
		case 0:
			Res = append(Res, matrix[up][left:right+1]...)
			up = up + 1
		case 1:
			for i := up; i <= down; i++ {
				Res = append(Res, matrix[i][right])
			}
			right = right - 1
		case 2:
			for j := right; j >= left; j-- {
				Res = append(Res, matrix[down][j])
			}
			down = down - 1
		case 3:
			for i := down; i >= up; i-- {
				Res = append(Res, matrix[i][left])
			}
			left = left + 1
		default:
			break
		}
	}
	return Res
}

func main() {
	exp2()
	fmt.Println("...")
}
