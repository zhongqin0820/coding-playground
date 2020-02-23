package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/long-pressed-name/
func isLongPressedName(name string, typed string) bool {
	res := true
	i := 0
	for j := 0; i < len(name) && j < len(typed); j++ {
		if name[i] == typed[j] {
			i++
			continue
		}
		if (j > 0 && typed[j] != typed[j-1]) || (j == 0) {
			res = false
		}
	}
	return res && i == len(name)
}

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		name   string
		typed  string
		output bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
		{"leelee", "eeelleeelee", false},
		{"pyplrz", "ppyypllr", false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isLongPressedName(ts.name, ts.typed))
		})
	}
}
