package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-168/problems/maximum-candies-you-can-get-from-boxes/
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	// TODO(zoking)[2019-12-22]: https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/discuss/457565/C%2B%2B-Simulation-using-BFS
	return 0
}

func TestMaxCandies(t *testing.T) {
	tests := []struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
		output         int
	}{
		{
			[]int{1, 0, 1, 0},
			[]int{7, 5, 4, 100},
			[][]int{[]int{}, []int{}, []int{1}, []int{}},
			[][]int{[]int{1, 2}, []int{3}, []int{}, []int{}},
			[]int{0},
			16,
		},
		{
			[]int{1, 0, 1, 0},
			[]int{1, 1, 1, 1, 1, 1},
			[][]int{[]int{1, 2, 3, 4, 5}, []int{}, []int{}, []int{}, []int{}, []int{}},
			[][]int{[]int{1, 2, 3, 4, 5}, []int{}, []int{}, []int{}, []int{}, []int{}},
			[]int{0},
			6,
		},
		{
			[]int{1, 1, 1},
			[]int{100, 1, 100},
			[][]int{[]int{}, []int{0, 2}, []int{}},
			[][]int{[]int{}, []int{}, []int{}},
			[]int{1},
			1,
		},
		{
			[]int{1},
			[]int{100},
			[][]int{[]int{}},
			[][]int{[]int{}},
			[]int{},
			0,
		},
		{
			[]int{1, 1, 1},
			[]int{2, 3, 2},
			[][]int{[]int{}, []int{}, []int{}},
			[][]int{[]int{}, []int{}, []int{}},
			[]int{2, 1, 0},
			7,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxCandies(ts.status, ts.candies, ts.keys, ts.containedBoxes, ts.initialBoxes))
		})
	}
}
