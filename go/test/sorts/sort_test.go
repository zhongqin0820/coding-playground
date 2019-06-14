package sorts

import (
	"testing"
)

// go test -v -test.run TestBubleSort
func TestBubleSort(t *testing.T) {
	s := []int{3, 6, 4, 2, 11, 10, 5}
	t.Run("BubleSort", func(t *testing.T) {
		BubleSort(s, false) //small -> big
		t.Log(s)
		BubleSort(s, true) //big -> small
		t.Log(s)
	})
}

// go test -v -test.run SelectSort
func TestSelectSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("SelectSort", func(t *testing.T) {
		SelectSort(s, false) //small -> big
		t.Log(s)
		SelectSort(s, true) //big -> small
		t.Log(s)
	})
}

// go test -v -test.run TestInsertSort
func TestInsertSort(t *testing.T) {
	s := []int{3, 6, 4, 2, 11, 10, 5}
	t.Run("InsertSortNavie", func(t *testing.T) {
		InsertSortNavie(s, false)
		t.Log(s)
		InsertSortNavie(s, true)
		t.Log(s)
	})

	s = []int{3, 6, 4, 2, 11, 10, 5}
	t.Run("InsertSort", func(t *testing.T) {
		InsertSort(s, false) //small -> big
		t.Log(s)
		InsertSort(s, true) //big -> small
		t.Log(s)
	})

	s = []int{3, 6, 4, 2, 11, 10, 5}
	t.Run("ShellSort", func(t *testing.T) {
		ShellSort(s, false)
		t.Log(s)
		ShellSort(s, true)
		t.Log(s)
	})
}

// go test -v -test.run MergeSort
func TestMergeSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	s1 := MergeSort(s, false) //small -> big
	t.Log(s1)
	s2 := MergeSort(s, true) //big -> small
	t.Log(s2)
}

// go test -v -test.run QuickSort
func TestQuickSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	QuickSort(s, false) //small -> big
	t.Log(s)
	QuickSort(s, true) //big -> small
	t.Log(s)
}

// go test -v -run=HeapSort
func TestHeapSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	HeapSort(s, false) //small -> big
	t.Log(s)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	HeapSort(s, true) //big -> small
	t.Log(s)
}
