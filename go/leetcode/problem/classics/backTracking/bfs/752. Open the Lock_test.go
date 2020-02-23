package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func openLock(deadends []string, target string) int {
	deads := make(map[string]bool, len(deadends))
	for _, s := range deadends {
		deads[s] = true
	}
	if deads["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	visited := make(map[string]bool, 10000)
	visited["0000"] = true
	res := 0
	d := [2]int{-1, 1}
	q := make([]string, 1, 10000)
	q[0] = "0000"

	for l := len(q); l > 0; l = len(q) {
		res++
		for i := 0; i < l; i++ {
			s := q[i]
			for j := 0; j < len(s); j++ {
				for _, dj := range d {
					t := []byte(s)
					t[j] = byte((int(t[j]-'0')+dj+10)%10 + '0')
					s_ := string(t)
					if s_ == target {
						return res
					}
					if !deads[s_] && !visited[s_] {
						q = append(q, s_)
					}
					visited[s_] = true
				}
			}
		}
		q = q[l:]
	}
	return -1
}

func TestOpenLock(t *testing.T) {
	tests := []struct {
		deadends []string
		target   string
		output   int
	}{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
		{[]string{"0000"}, "8888", -1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, openLock(ts.deadends, ts.target))
		})
	}
}
