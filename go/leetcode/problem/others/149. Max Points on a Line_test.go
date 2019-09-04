package problem

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-points-on-a-line/
func maxPoints(points [][]int) int {
	// n * 2
	// edge case
	if points == nil || len(points) == 0 || points[0] == nil || len(points[0]) == 0 {
		return 0
	}
	n := len(points)
	if n == 1 {
		return n
	}
	// 计算有多少个不同的X坐标, value[0]是计数器, value=append(value,points[i][1])
	diffMap := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		if _, ok := diffMap[points[i][0]]; ok {
			diffMap[points[i][0]][0]++
			diffMap[points[i][0]] = append(diffMap[points[i][0]], points[i][1])
		} else {
			diffMap[points[i][0]] = []int{1, points[i][1]}
		}
	}
	size := len(diffMap)
	if size == 1 {
		// 所有的点都在同一条X坐标直线上
		return n
	}
	fmt.Println(diffMap)
	// 去重
	var max int
	var x int
	var ys []int
	for x, ys = range diffMap {
		tempN := len(ys)
		temp := make(map[int]struct{}, tempN)
		// 去重ys
		for i, y := range ys {
			if i == 0 {
				continue
			}
			temp[y] = struct{}{}
		}
		// 清空x列表
		diffMap[x] = []int{}
		for y, _ := range temp {
			diffMap[x] = append(diffMap[x], y)
		}
		// 获取长度append到diffMap[x]头部
		tempN = len(diffMap[x])
		diffMap[x] = append([]int{tempN}, diffMap[x]...)
		if max < tempN {
			max = tempN
		}
	}
	fmt.Println(diffMap)
	// 计算p1,p2之间的k
	// 需要计算(n-1)*n/2次
	return max
}

func helper149(p1, p2 []int) string {
	return strconv.FormatFloat(float64((p1[1]-p2[2])/(p1[0]-p2[0])), 'f', -1, 64)
}

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		output int
	}{
		{[][]int{{1, 1}, {1, 1}, {3, 3}}, 3},
		// {[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxPoints(ts.points))
		})
	}
}
