package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/queens-that-can-attack-the-king/
// ref: https://leetcode.com/problems/queens-that-can-attack-the-king/discuss/403669/Python-Check-8-steps-in-8-Directions
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	N := 8
	res := make([][]int, 0, N)
	// 哈希表记录queen的位置
	m := make(map[string]bool, len(queens))
	for _, queen := range queens {
		m[fmt.Sprintf("%d,%d", queen[0], queen[1])] = true
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// i,j 对应8个方向
			// d表示深度
			for d := 1; d < N; d++ {
				x, y := king[0]+i*d, king[1]+j*d
				if _, ok := m[fmt.Sprintf("%d,%d", x, y)]; ok {
					res = append(res, []int{x, y})
					break
				}
			}
		}
	}
	return res
}

func TestQueensAttacktheKing(t *testing.T) {
	tests := []struct {
		queens [][]int
		king   []int
		output [][]int
	}{
		{
			[][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
			[]int{0, 0},
			[][]int{{0, 1}, {1, 0}, {3, 3}},
		},
		{
			[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}},
			[]int{3, 3},
			[][]int{{2, 2}, {3, 4}, {4, 4}},
		},
		{
			[][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}},
			[]int{3, 4},
			[][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}},
		},
		{
			[][]int{{7, 3}, {4, 7}, {1, 3}, {3, 0}, {6, 2}, {1, 6}, {5, 1}, {3, 7}, {7, 2}, {4, 0}, {3, 3}, {5, 5}, {4, 4}, {6, 3}, {1, 5}, {4, 1}, {1, 1}, {6, 4}, {2, 6}, {0, 0}, {7, 1}, {4, 5}, {6, 0}, {7, 5}, {0, 5}, {6, 5}, {3, 5}, {2, 7}, {0, 1}, {5, 3}, {7, 0}, {4, 6}, {5, 7}, {7, 4}, {0, 6}, {2, 0}, {4, 3}, {3, 4}},
			[]int{0, 2},
			[][]int{{0, 1}, {0, 5}, {1, 1}, {6, 2}, {1, 3}},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, queensAttacktheKing(ts.queens, ts.king))
		})
	}
}
