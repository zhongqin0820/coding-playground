package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/course-schedule/description/
func canFinish(numCourses int, prerequisites [][]int) bool {
	next := helper207buildNext(numCourses, prerequisites)
	pres := helper207buildPres(next)
	return helper207check(next, pres, numCourses)
}

// 课程对应的先修课
func helper207buildNext(n int, prerequisites [][]int) [][]int {
	next := make([][]int, n)
	for _, pair := range prerequisites {
		next[pair[1]] = append(next[pair[1]], pair[0])
	}
	return next
}

//对应的先修课的个数
func helper207buildPres(next [][]int) []int {
	pres := make([]int, len(next))
	for _, neighbours := range next {
		for _, n := range neighbours {
			pres[n]++
		}
	}
	return pres
}

func helper207check(next [][]int, pres []int, numCourses int) bool {
	var i, j int
	for i = 0; i < numCourses; i++ {
		for j = 0; j < numCourses; j++ {
			if pres[j] == 0 {
				break
			}
		}
		// 每个课程都需要先修课
		// 出现了环路
		if j >= numCourses {
			return false
		}
		// 修改 pres[j] 为负数
		// 避免重修
		pres[j] = -1
		// 完成 j 课程后
		// j 的后续课程的，先修课程数量都可以 -1
		for _, c := range next[j] {
			pres[c]--
		}
	}

	return true
}

//解决方案2
func canFinishI(numCourses int, prerequisites [][]int) bool {
	// 初始化字典：key是前置课程，列表是需要选修该前置课程的课程
	prevs := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		prevs[prerequisites[i][0]] = append(prevs[prerequisites[i][0]], prerequisites[i][1])
	}
	//每个课的所有先修课程都必须是能修完的，这个课才是能修完的
	for prev, _ := range prevs {
		var finded []int
		if !helper207(prev, &prevs, finded) {
			return false
		}
	}
	return true
}

func helper207(x int, course *map[int][]int, finded []int) bool {
	finded = append(finded, x)
	if v, ok := (*course)[x]; ok {
		for k := 0; k < len(v); k++ {
			for _, d := range finded {
				if d == v[k] {
					return false
				}
			}
			if !helper207(v[k], course, finded) {
				return false
			}
		}
	}
	//课程成功修完的这条路线上的路程都能被修完，把它们都从map里删掉，避免后面重复检查
	delete(*course, x)
	return true
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		output        bool
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
			true,
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
			false,
		},
		{
			2,
			[][]int{
				[]int{0, 1},
			},
			true,
		},
		{
			2,
			[][]int{
				[]int{1, 0},
			},
			true,
		},
		{
			2,
			[][]int{
				[]int{1, 0},
				[]int{0, 1},
			},
			false,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, canFinish(ts.numCourses, ts.prerequisites))
			ast.Equal(ts.output, canFinishI(ts.numCourses, ts.prerequisites))
		})
	}
}
