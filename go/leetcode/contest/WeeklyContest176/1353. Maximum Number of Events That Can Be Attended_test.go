package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
func maxEvents(events [][]int) int {
	// ref: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/discuss/510246/JavaScript-Greedy
	// ref: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/discuss/510263/JavaC%2B%2BPython-Priority-Queue
	m := make(map[int]struct{}, 100000)
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] != events[j][1] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	res := 0
	for i := 0; i < n; i++ {
		for d := events[i][0]; d <= events[i][1]; d++ {
			if _, ok := m[d]; !ok {
				m[d] = struct{}{}
				res++
				break
			}
		}
	}
	return res
}

func TestMaxEvents(t *testing.T) {
	tests := []struct {
		events [][]int
		output int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 3},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}, 4},
		{[][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}, 4},
		{[][]int{{1, 5}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}, 5},
		{[][]int{{1, 100000}}, 1},
		{[][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}, 7},
		{[][]int{{1, 3}, {1, 3}, {1, 3}, {3, 4}}, 4},
		{[][]int{{26, 27}, {24, 26}, {1, 4}, {1, 2}, {3, 6}, {2, 6}, {9, 13}, {6, 6}, {25, 27}, {22, 25}, {20, 24}, {8, 8}, {27, 27}, {30, 32}}, 14},
		{[][]int{{6, 6}, {3, 4}, {22, 26}, {29, 32}, {8, 10}, {8, 9}, {30, 30}, {19, 21}, {30, 34}, {20, 20}, {29, 32}, {4, 5}, {16, 17}, {3, 3}, {14, 16}, {9, 10}, {2, 5}, {7, 11}, {3, 3}, {18, 20}, {26, 28}, {15, 19}, {26, 27}, {22, 22}, {2, 3}, {16, 20}, {2, 3}, {23, 27}, {25, 28}, {17, 20}}, 27},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxEvents(ts.events))
		})
	}
}
