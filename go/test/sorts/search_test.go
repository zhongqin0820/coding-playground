package sorts

import (
	"testing"
)

// go test -v -run=TestInterpolationSearch
func TestInterpolationSearch(t *testing.T) {
	a := []int{3, 5, 1, 6, 7, 9, 10}
	QuickSort(a, false)

	t.Run("BinarySearch 1", func(t *testing.T) {
		target := 5
		if x := BinarySearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("BinarySearch 2", func(t *testing.T) {
		target := -4
		if x := BinarySearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})

	t.Run("InterpolationSearch 1", func(t *testing.T) {
		target := 5
		if x := InterpolationSearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("InterpolationSearch 2", func(t *testing.T) {
		target := -4
		if x := InterpolationSearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})

	t.Run("FibSearch 1", func(t *testing.T) {
		target := 5
		if x := FibSearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("FibSearch 2", func(t *testing.T) {
		target := -4
		if x := FibSearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})

}

// go test -v -run=TestOtherSearch
// HashSearch, SeqSearch
func TestOtherSearch(t *testing.T) {
	a := []int{3, 5, 1, 6, 7, 9, 10}
	QuickSort(a, false)
	t.Run("HashSearch 1", func(t *testing.T) {
		target := 5
		if x := HashSearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("HashSearch 2", func(t *testing.T) {
		target := -4
		if x := HashSearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("SeqSearch 1", func(t *testing.T) {
		target := 5
		if x := SeqSearch(a, target); x == -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
	t.Run("SeqSearch 2", func(t *testing.T) {
		target := -4
		if x := SeqSearch(a, target); x != -1 {
			t.Error("Error")
		} else {
			t.Log(x, target)
		}
	})
}
