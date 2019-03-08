package sortAlgorithm

import (
	"testing"
)

// go test -v -test.run TestBinarySearch
func TestBinarySearch(t *testing.T) {
	a := []int{3, 5, 1, 6, 7, 9, 10}
	a = QuickSort(a, false)
	// t.Log(a)
	target := 5
	if x := BinarySearch(a, target); x == -1 {
		t.Error("Error")
	} else {
		t.Log(x)
	}
	target = -4
	if x := BinarySearch(a, target); x != -1 {
		t.Error("Error")
	} else {
		t.Log(target)
	}

}
