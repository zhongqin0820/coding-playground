package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type rest struct {
	id  int
	rat int
}

type rests []rest

func (r rests) Len() int { return len(r) }
func (r rests) Less(i, j int) bool {
	if r[i].rat == r[j].rat {
		return r[i].id > r[j].id
	}
	return r[i].rat > r[j].rat
}
func (r rests) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	n := len(restaurants)
	rs := make(rests, 0, n)
	for _, r := range restaurants {
		if (r[2] == veganFriendly || veganFriendly == 0) && r[3] <= maxPrice && r[4] <= maxDistance {
			rs = append(rs, rest{
				id:  r[0],
				rat: r[1],
			})
		}
	}
	sort.Sort(rs)
	res := make([]int, len(rs))
	for i, r := range rs {
		res[i] = r.id
	}
	return res
}

func TestFilterRestaurants(t *testing.T) {
	tests := []struct {
		restaurants   [][]int
		veganFriendly int
		maxPrice      int
		maxDistance   int
		output        []int
	}{
		{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 1, 50, 10, []int{3, 1, 5}},
		{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 50, 10, []int{4, 3, 2, 1, 5}},
		{[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 30, 3, []int{4, 5}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, filterRestaurants(ts.restaurants, ts.veganFriendly, ts.maxPrice, ts.maxDistance))
		})
	}
}
