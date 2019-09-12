package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/isomorphic-strings/description/

// Given two strings s and t, determine if they are isomorphic.

// Two strings are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1, m2 := [256]int{}, [256]int{}
	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
}

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t   string
		output bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, isIsomorphic(ts.s, ts.t))
		})
	}
}
