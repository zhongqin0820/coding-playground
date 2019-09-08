package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})
	return letters[i%n]
}

func naiveNextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l < n {
		return letters[l]
	}
	return letters[0]
}

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		output  byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'k', 'c'},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, nextGreatestLetter(ts.letters, ts.target))
			ast.Equal(ts.output, naiveNextGreatestLetter(ts.letters, ts.target))
		})
	}
}
