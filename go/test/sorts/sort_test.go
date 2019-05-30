package sorts

import (
	"testing"
)

// go test -v -test.run TestBubleSort
func TestBubleSort(t *testing.T) {
	s := []int{3, 6, 4, 2, 11, 10, 5}
	s1 := BubleSort(s, false) //small -> big
	t.Log(s1)
	s2 := BubleSort(s, true) //big -> small
	t.Log(s2)
}

// go test -v -test.run TestInsertSort
func TestInsertSort(t *testing.T) {
	s := []int{3, 6, 4, 2, 11, 10, 5}
	s1 := InsertSort(s, false) //small -> big
	t.Log(s1)
	s2 := InsertSort(s, true) //big -> small
	t.Log(s2)
}

// go test -v -test.run SelectSort
func TestSelectSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	s1 := InsertSort(s, false) //small -> big
	t.Log(s1)
	s2 := InsertSort(s, true) //big -> small
	t.Log(s2)
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
	s1 := QuickSort(s, false) //small -> big
	t.Log(s1)
	s2 := QuickSort(s, true) //big -> small
	t.Log(s2)
}

// go test -v -test.run InsertReview
func TestInsertReview(t *testing.T) {
	s := []int{-1, 3, 2, -6}
	s1 := InsertReview(s)
	t.Log(s1)
}

// go test -v -test.run SelectReview
func TestSelectReview(t *testing.T) {
	s := []int{-1, 3, 2, -5}
	s1 := SelectReview(s)
	t.Log(s1)
}
