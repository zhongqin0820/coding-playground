package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
type Points [][]int

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findMinArrowShots(points [][]int) int {
	n, res := len(points), 0
	if n == 0 {
		return res
	}
	// 按照末尾值升序，小值在前
	sort.Sort(Points(points))
	end := points[0][1]
	for i := 1; i < n; i++ {
		// 当前范围的起始<=前一个范围的结束
		// 符合shoot的范围
		if points[i][0] <= end {
			continue
		}
		// shoot++同时更新结束值
		res++
		end = points[i][1]
	}
	res++
	return res
}

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		points [][]int
		output int
	}{
		{[][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}, 2},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findMinArrowShots(ts.points))
		})
	}
}
