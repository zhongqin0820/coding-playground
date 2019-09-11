package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/is-graph-bipartite/description/
func isBipartite(graph [][]int) bool {
	n := len(graph)
	// painted[i] == 1  表示 node i 在 白色 集合中
	// painted[i] == -1 表示 node i 在 黑色 集合中
	// painted[i] == 0  表示 node i 还没有被标记集合
	painted := make([]int, n)

	for i := 0; i < n; i++ {
		// painted[i]== 0 说明，所有与 node i 相互联通的点，都没有被检查过
		if painted[i] == 0 && !helper785(i, 1, painted, graph) {
			return false
		}
	}
	return true
}

func helper785(node, color int, painted []int, graph [][]int) bool {
	// 如果 node 已经被检查过了
	// 可以直接核对颜色
	if painted[node] != 0 {
		return painted[node] == color
	}
	// 按照要求，对 node 进行归类
	painted[node] = color
	// 对与 node 连接的点，也进行同样的检查
	for _, i := range graph[node] {
		if !helper785(i, -color, painted, graph) {
			return false
		}
	}
	return true
}

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		graph  [][]int
		output bool
	}{
		{
			[][]int{{}, {2}, {1}, {4}, {3}}, // 图是两根分开的边
			true,
		},
		{
			[][]int{{2, 4}, {2, 3, 4}, {0, 1}, {1}, {0, 1}, {7}, {9}, {5}, {}, {6}, {12, 14}, {}, {10}, {}, {10}, {19}, {18}, {}, {16}, {15}, {23}, {23}, {}, {20, 21}, {}, {}, {27}, {26}, {}, {}, {34}, {33, 34}, {}, {31}, {30, 31}, {38, 39}, {37, 38, 39}, {36}, {35, 36}, {35, 36}, {43}, {}, {}, {40}, {}, {49}, {47, 48, 49}, {46, 48, 49}, {46, 47, 49}, {45, 46, 47, 48}},
			false,
		},

		{
			[][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}},
			false,
		},

		{
			[][]int{{}},
			true,
		},

		{
			[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			true,
		},

		{
			[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			false,
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isBipartite(ts.graph))
		})
	}
}
