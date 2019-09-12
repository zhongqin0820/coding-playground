package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/queue-reconstruction-by-height/description/
type People [][]int

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 身高降序，人数升序
func (p People) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] > p[j][0]
}

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))
	sort.Sort(People(people))
	insert := func(idx int, person []int) {
		res = append(res, person)
		if len(res)-1 == idx {
			return
		}
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}
	return res
}

func TestReconstructQueue(t *testing.T) {
	tests := []struct {
		people [][]int
		output [][]int
	}{
		{
			[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, reconstructQueue(ts.people))
		})
	}
}
