package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numOfSubarrays(arr []int, k int, threshold int) int {
	SUM := k * threshold
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	res := 0
	for i, j := 0, k; j <= len(arr); i, j = i+1, j+1 {
		if sum >= SUM {
			res++
		}
		if j < len(arr) {
			sum = sum - arr[i] + arr[j]
		}
	}
	return res
}

func TestNumOfSubarrays(t *testing.T) {
	tests := []struct {
		arr       []int
		k         int
		threshold int
		output    int
	}{
		{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4, 3},
		{[]int{1, 1, 1, 1, 1}, 1, 0, 5},
		{[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5, 6},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 7, 7, 1},
		{[]int{4, 4, 4, 4}, 4, 1, 1},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numOfSubarrays(ts.arr, ts.k, ts.threshold))
		})
	}
}
