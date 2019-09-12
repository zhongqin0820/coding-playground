package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	next, pre := helper210build(numCourses, prerequisites)
	return helper210search(next, pre)
}

func helper210build(num int, requires [][]int) (next [][]int, pre []int) {
	next = make([][]int, num)
	pre = make([]int, num)
	for _, r := range requires {
		next[r[1]] = append(next[r[1]], r[0])
		pre[r[0]]++
	}
	return
}

func helper210search(next [][]int, pre []int) []int {
	n, res := len(pre), make([]int, len(pre))
	for i, j := 0, 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if pre[j] == 0 {
				break
			}
		}
		if j == n {
			return nil
		}
		pre[j] = -1
		for _, c := range next[j] {
			pre[c]--
		}
		res[i] = j
	}
	return res
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		output        []int
	}{
		{
			4,
			[][]int{
				[]int{1, 0},
				[]int{2, 1},
				[]int{2, 0},
				[]int{3, 0},
				[]int{3, 1},
				[]int{3, 2},
			},
			[]int{0, 1, 2, 3},
		},
		{
			4,
			[][]int{
				[]int{1, 0},
				[]int{2, 1},
				[]int{2, 0},
				[]int{3, 0},
				[]int{3, 1},
				[]int{3, 2},
				[]int{2, 3},
			},
			nil,
		},
		{
			2,
			[][]int{
				[]int{0, 1},
			},
			[]int{1, 0},
		},
		{
			2,
			[][]int{
				[]int{1, 0},
			},
			[]int{0, 1},
		},
		{
			2,
			[][]int{
				[]int{1, 0},
				[]int{0, 1},
			},
			nil,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findOrder(ts.numCourses, ts.prerequisites))
		})
	}
}
