package problem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/to-lower-case/

// Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

func toLowerCase(str string) string {
	bs := []byte(str)
	for i := 0; i < len(bs); i++ {
		if bs[i] >= 'A' && bs[i] <= 'Z' {
			bs[i] = bs[i] - 'A' + 'a'
		}
	}
	return string(bs)
}

func toLowerCaseAdv(str string) string {
	return strings.ToLower(str)
}

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"Hello", "hello"},
		{"hello", "hello"},
		{"LOVELY", "lovely"},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, toLowerCase(ts.input))
		})
	}
}

func BenchmarkToLowerCase(b *testing.B) {
	b.StartTimer()
	b.Run("Naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			toLowerCase("LOVEly")
		}
	})
	b.ResetTimer()
	b.StartTimer()
	b.Run("Adv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			toLowerCaseAdv("LOVEly")
		}
	})
}
