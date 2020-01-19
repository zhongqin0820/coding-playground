package contest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-172/problems/print-words-vertically/
func printVertically(s string) []string {
	words := strings.Split(s, " ")
	max_wz := 0
	for _, word := range words {
		if max_wz < len(word) {
			max_wz = len(word)
		}
	}
	res := make([]string, max_wz)
	for i := 0; i < max_wz; i++ {
		for _, word := range words {
			if i >= len(word) {
				res[i] += " "
			} else {
				res[i] += string(word[i])
			}
		}
		res[i] = strings.TrimRight(res[i], " ")
	}
	return res
}

func TestPrintVertically(t *testing.T) {
	tests := []struct {
		s      string
		output []string
	}{
		{"HOW ARE YOU", []string{"HAY", "ORO", "WEU"}},
		{"TO BE OR NOT TO BE", []string{"TBONTB", "OEROOE", "   T"}},
		{"CONTEST IS COMING", []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, printVertically(ts.s))
		})
	}
}
