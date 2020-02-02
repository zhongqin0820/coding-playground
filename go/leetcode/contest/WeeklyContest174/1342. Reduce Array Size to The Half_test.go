package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct {
	key   int
	value int
}

type counters []pair

func minSetSize(arr []int) int {
	n := len(arr)
	k := n >> 1
	counter := make(map[int]int, n)
	for _, k := range arr {
		counter[k]++
	}

	c := make(counters, 0, len(counter))
	for k, v := range counter {
		c = append(c, pair{
			key:   k,
			value: v,
		})
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i].value >= c[j].value
	})
	res := 0
	for _, p := range c {
		if k <= 0 {
			break
		}
		k -= p.value
		res++
	}
	return res
}

func TestMinSetSize(t *testing.T) {
	tests := []struct {
		arr    []int
		output int
	}{
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 9}, 1},
		{[]int{1000, 1000, 3, 7}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, minSetSize(ts.arr))
		})
	}
}
