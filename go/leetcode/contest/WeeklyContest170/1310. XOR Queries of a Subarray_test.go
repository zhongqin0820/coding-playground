package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-170/problems/xor-queries-of-a-subarray/
// challenge:
//      1 <= arr.length <= 3 * 10^4
//      1 <= queries.length <= 3 * 10^4
func xorQueries(arr []int, queries [][]int) []int {
	n := len(queries)
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = arr[queries[i][0]]
		for j := queries[i][0] + 1; j <= queries[i][1]; j++ {
			res[i] ^= arr[j]
		}
	}
	return res
}

// https://leetcode.com/problems/xor-queries-of-a-subarray/discuss/470787/JavaC%2B%2BPython-Straight-Forward-Solution
func xorQueriesII(arr []int, queries [][]int) []int {
	m := len(arr)
	for i := 1; i < m; i++ {
		arr[i] ^= arr[i-1]
	}

	n := len(queries)
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		if queries[i][0] == 0 {
			res[i] = arr[queries[i][1]]
		} else {
			res[i] = arr[queries[i][0]-1] ^ arr[queries[i][1]]
		}
	}
	return res
}

func TestXorQueries(t *testing.T) {
	tests := []struct {
		arr     []int
		queries [][]int
		output  []int
	}{
		{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}, []int{2, 7, 14, 8}},
		{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}, []int{8, 0, 4, 4}},
		{[]int{16}, [][]int{{0, 0}, {0, 0}, {0, 0}}, []int{16, 16, 16}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, xorQueries(ts.arr, ts.queries))
			ast.Equal(ts.output, xorQueriesII(ts.arr, ts.queries))
		})
	}
}
