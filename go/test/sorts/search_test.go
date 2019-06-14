package sorts

import (
	"testing"
)

// go test -v -test.run TestSearch
func TestSearch(t *testing.T) {
	a := []int{3, 5, 1, 6, 7, 9, 10}
	QuickSort(a, false)

	t.Run("BinarySearch 1", func(t *testing.T) {
		target := 5
		if x := BinarySearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x)
		}
	})

	t.Run("BinarySearch 2", func(t *testing.T) {
		target := -4
		if x := BinarySearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(target)
		}
	})

	t.Run("InterpolationSearch 1", func(t *testing.T) {
		target := 5
		if x := InterpolationSearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(target)
		}
	})

	t.Run("InterpolationSearch 2", func(t *testing.T) {
		target := -4
		if x := InterpolationSearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(target)
		}
	})
}
