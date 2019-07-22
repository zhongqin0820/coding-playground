package sorts

// need to assign the mergesort result to a slice and rearrange the two parts from the assigned slice
func MergeSort(s []int, flag bool) []int {
	if Defensive(s) {
		return s
	}
	n := len(s)
	m := n / 2
	a := MergeSort(s[:m], flag)
	b := MergeSort(s[m:], flag)
	return Merge(a, b, flag)
}

func Merge(a, b []int, flag bool) (ret []int) {
	var i, j int
	if !flag { //small -> big
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] < b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
	} else { //big -> small
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] > b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
	}
	ret = append(ret, a[i:]...)
	ret = append(ret, b[j:]...)
	return ret
}

// the order of s[k] and v decide the sorted order of the array
func QuickSort(s []int, flag bool) {
	if Defensive(s) {
		return
	}
	// v is the value of pivot
	// k is the current index
	var i, j, v, k = 0, len(s) - 1, s[0], 1
	for i < j {
		if !flag { //small -> big
			if s[k] > v {
				s[k], s[j] = s[j], s[k]
				j--
			} else {
				s[i], s[k] = s[k], s[i]
				i++
				k++
			}
		} else { //big->small
			if s[k] < v {
				s[k], s[j] = s[j], s[k]
				j--
			} else {
				s[i], s[k] = s[k], s[i]
				i++
				k++
			}
		}
	}
	QuickSort(s[:i], flag)
	QuickSort(s[i+1:], flag)
}

// heap sort = build heap + adjust heap

// adjust the value node[i] with its sons
func HeapAdjust(a []int, i, n int, flag bool) {
	var ichild int
	for ; 2*i+1 < n; i = ichild {
		ichild = 2*i + 1
		if !flag { //max heap: sort order small -> big
			if ichild < n-1 && a[ichild+1] > a[ichild] {
				ichild++
			}
			if a[i] < a[ichild] {
				a[i], a[ichild] = a[ichild], a[i]
			} else {
				break
			}
		} else { //min heap: sort order big -> small
			if ichild < n-1 && a[ichild+1] < a[ichild] {
				ichild++
			}
			if a[i] > a[ichild] {
				a[i], a[ichild] = a[ichild], a[i]
			} else {
				break
			}
		}

	}
}

func HeapSort(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	n := len(a)
	// build heap
	for i := n / 2; i >= 0; i-- {
		// adjust heap
		// log.Println(a)
		HeapAdjust(a, i, n, flag)
	}
	// sort out
	for i := n - 1; i > 0; i-- {
		// swap
		a[0], a[i] = a[i], a[0]
		// adjust the remaining
		HeapAdjust(a, 0, i, flag)
	}
}
