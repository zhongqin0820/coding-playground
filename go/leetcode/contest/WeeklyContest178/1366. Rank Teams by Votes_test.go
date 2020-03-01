package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type s struct {
	t   byte
	cnt []int
}

func rankTeams(votes []string) string {
	m := make(map[byte][]int, 26)
	for _, vote := range votes {
		for i := range vote {
			if _, ok := m[vote[i]]; !ok {
				m[vote[i]] = make([]int, len(vote))
			}
			m[vote[i]][i]++
		}
	}
	temp := make([]s, 0, len(m))
	for t, cnt := range m {
		temp = append(temp, s{
			t:   t,
			cnt: cnt,
		})
	}
	sort.Slice(temp, func(i, j int) bool {
		for k, cnt := range temp[i].cnt {
			if cnt == temp[j].cnt[k] {
				continue
			}
			return cnt > temp[j].cnt[k]
		}
		return temp[i].t < temp[j].t
	})
	res := make([]byte, len(m))
	for i := range temp {
		res[i] = temp[i].t
	}
	return string(res)
}
func TestRankTeams(t *testing.T) {
	tests := []struct {
		votes  []string
		output string
	}{
		{[]string{"ABC", "ACB", "ABC", "ACB", "ACB"}, "ACB"},
		{[]string{"WXYZ", "XYZW"}, "XWYZ"},
		{[]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}, "ZMNAGUEDSJYLBOPHRQICWFXTVK"},
		{[]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}, "ABC"},
		{[]string{"M", "M", "M", "M"}, "M"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, rankTeams(ts.votes))
		})
	}
}
