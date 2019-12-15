package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/contest/weekly-contest-167/problems/sequential-digits/
func sequentialDigits(low int, high int) []int {
	res := make([]int, 0, 2048)
	for i := 1; i < 10; i++ {
		num := 0
		for j := i; j < 10; j++ {
			num = num*10 + j
			if num > high {
				break
			}
			if num >= low {
				res = append(res, num)
			}
		}
	}
	sort.Ints(res)
	return res
}

func TestSequentialDigits(t *testing.T) {
	tests := []struct {
		low    int
		high   int
		output []int
	}{
		{100, 300, []int{123, 234}},
		{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
		{10, 1000000000, []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, sequentialDigits(ts.low, ts.high))
		})
	}
}
