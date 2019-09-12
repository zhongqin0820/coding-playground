package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/
func reverseVowels(s string) string {
	var m map[byte]struct{} = map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}, 'A': struct{}{}, 'E': struct{}{}, 'I': struct{}{}, 'O': struct{}{}, 'U': struct{}{}}
	i, j := 0, len(s)-1
	sb := []byte(s)
	for i < j {
		_, okI := m[sb[i]]
		_, okJ := m[sb[j]]
		if okI && okJ {
			sb[i], sb[j] = sb[j], sb[i]
			i++
			j--
		}
		if !okI {
			i++
		}
		if !okJ {
			j--
		}
	}
	return string(sb)
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, reverseVowels(ts.s))
		})
	}
}
