package problemSet

import (
	"fmt"
	"testing"
)

func TeststrStr(t *testing.T) {
	tm := []struct {
		hay      string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}
	for i, c := range tm {
		t.Run(fmt.Sprintf("Test%d\n", i), func(t *testing.T) {
			res := strStr(c.hay, c.needle)
			if res != c.expected {
				t.Errorf("want %d, but got %d\n", c.expected, res)
			}
		})
	}
}
