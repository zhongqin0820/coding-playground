package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if target-numbers[i] == numbers[j] {
			break
		} else if target-numbers[i] > numbers[j] {
			i++
		} else {
			j--
		}
	}
	return []int{i + 1, j + 1}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		output  []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, twoSum(ts.numbers, ts.target))
		})
	}
}
