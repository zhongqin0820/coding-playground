package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-chunks-to-make-sorted/description/

// Given an array arr that is a permutation of [0, 1, ..., arr.length - 1], we split the array into some number of "chunks" (partitions), and individually sort each chunk.  After concatenating them, the result equals the sorted array.

// What is the most number of chunks we could have made?
// Note:
// 1. arr will have length in range [1, 10].
// 1. arr[i] will be a permutation of [0, 1, ..., arr.length - 1].

func maxChunksToSorted(arr []int) int {
	// 假设某个 chunk 为 arr[j:k]，则对于
	// j <=     i  < k ，必有
	// j <= arr[i] < k
	// lastIdx 是 下一个要切下来的 chunk 中最后一个元素的索引号
	lastIdx, res, n := 0, 0, len(arr)
	for i := 0; i < n; i++ {
		if lastIdx < arr[i] {
			// lastIdx == arr[i]，才能满足前面提到的要求
			lastIdx = arr[i]
			continue
		}
		// 此时可以切下一刀
		if i == lastIdx {
			res++
			lastIdx++
		}
	}
	return res
}

func TestMaxChunksToSorted(t *testing.T) {
	tests := []struct {
		arr    []int
		output int
	}{
		{[]int{2, 1, 0, 5, 4, 3, 6, 9, 8, 7}, 4},
		{[]int{4, 3, 2, 1, 0}, 1},
		{[]int{1, 0, 2, 3, 4}, 4},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, maxChunksToSorted(ts.arr))
		})
	}
}
