package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkIfExist(arr []int) bool {
	size := len(arr)
	sort.Ints(arr)
	m := make(map[int]int, size)
	// 大于0
	i := size - 1
	for ; i >= 0; i-- {
		if arr[i] < 0 {
			break
		}
		double := arr[i] * 2
		if _, ok := m[double]; ok {
			return true
		} else {
			m[arr[i]] = i
		}
	}
	// 小于0
	for j := 0; j <= i; j++ {
		double := arr[j] * 2
		if _, ok := m[double]; ok {
			return true
		} else {
			m[arr[j]] = j
		}
	}
	return false
}

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		arr    []int
		output bool
	}{
		{[]int{10, 2, 5, 3}, true},
		{[]int{7, 1, 14, 11}, true},
		{[]int{3, 1, 7, 11}, false},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, checkIfExist(ts.arr))
		})
	}
}
