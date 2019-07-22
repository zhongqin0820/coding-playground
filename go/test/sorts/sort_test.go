package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -test.run TestBubleSort
func TestBubleSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("BubleSort", func(t *testing.T) {
		ast := assert.New(t)
		BubleSort(s, false) //small -> big
		ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
		s = []int{3, 1, 4, 2, 11, 10, 5}
		BubleSort(s, true) //big -> small
		ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
	})
}

// go test -v -test.run SelectSort
func TestSelectSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("SelectSort", func(t *testing.T) {
		ast := assert.New(t)
		SelectSort(s, false) //small -> big
		ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
		s = []int{3, 1, 4, 2, 11, 10, 5}
		SelectSort(s, true) //big -> small
		ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
	})
}

// go test -v -test.run TestInsertSort
func TestInsertSort(t *testing.T) {
	s := []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("InsertSortNavie", func(t *testing.T) {
		ast := assert.New(t)
		InsertSortNavie(s, false)
		ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
		s = []int{3, 1, 4, 2, 11, 10, 5}
		InsertSortNavie(s, true)
		ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
	})

	s = []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("InsertSort", func(t *testing.T) {
		ast := assert.New(t)
		InsertSort(s, false) //small -> big
		ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
		s = []int{3, 1, 4, 2, 11, 10, 5}
		InsertSort(s, true) //big -> small
		ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
	})

	s = []int{3, 1, 4, 2, 11, 10, 5}
	t.Run("ShellSort", func(t *testing.T) {
		ast := assert.New(t)
		ShellSort(s, false)
		ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
		ShellSort(s, true)
		ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
	})
}

// go test -v -test.run MergeSort
func TestMergeSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	s1 := MergeSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s1)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	s2 := MergeSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s2)
}

// go test -v -test.run QuickSort
func TestQuickSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	QuickSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
	QuickSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
}

// go test -v -run=HeapSort
func TestHeapSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	HeapSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	HeapSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
}

// go test -v -run=CountingSort
func TestCountingSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	CountingSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	CountingSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
}

// go test -v -run=BucketSort
func TestBucketSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	BucketSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	BucketSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
}

// go test -v -run=RadixSort
func TestRadixSort(t *testing.T) {
	ast := assert.New(t)
	s := []int{3, 1, 4, 2, 11, 10, 5}
	RadixSort(s, false) //small -> big
	ast.Equal([]int{1, 2, 3, 4, 5, 10, 11}, s)
	s = []int{3, 1, 4, 2, 11, 10, 5}
	RadixSort(s, true) //big -> small
	ast.Equal([]int{11, 10, 5, 4, 3, 2, 1}, s)
}

// go test -v -run=^$ -bench=BenchmarkNSquare
func BenchmarkNSquare(b *testing.B) {
	b.ResetTimer()
	b.Run("BubleSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			BubleSort(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("SelectSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			SelectSort(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("InsertSortNavie", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			InsertSortNavie(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("InsertSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			InsertSort(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("ShellSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			ShellSort(s, false) //small -> big
		}
	})
}

// go test -v -run=^$ -bench=BenchmarkLogN
func BenchmarkLogN(b *testing.B) {
	b.ResetTimer()
	b.Run("MergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			MergeSort(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("QuickSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			QuickSort(s, false) //small -> big
		}
	})
	b.ResetTimer()
	b.Run("HeapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			HeapSort(s, false) //small -> big
		}
	})
}

// go test -v -run=^$ -bench=BenchmarkSortOthers
func BenchmarkSortOthers(b *testing.B) {
	b.ResetTimer()
	b.Run("CountingSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			CountingSort(s, false) //small -> big
		}
	})

	b.ResetTimer()
	b.Run("BucketSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			BucketSort(s, false) //small -> big
		}
	})

	b.ResetTimer()
	b.Run("RadixSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{3, 1, 4, 2, 11, 10, 5}
			RadixSort(s, false) //small -> big
		}
	})
}
