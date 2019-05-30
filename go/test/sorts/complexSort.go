package sorts

import (
	"log"
)

// need to assign the mergesort result to a slice and rearrange the two parts from the assigned slice
func MergeSort(s []int, flag bool) []int {
	Len := len(s)
	if Len <= 1 {
		return s
	}
	Mid := Len / 2
	a := MergeSort(s[:Mid], flag)
	b := MergeSort(s[Mid:], flag)
	return Merge(a, b, flag)
}

func Merge(a, b []int, flag bool) (ret []int) {
	var i, j int
	if !flag {
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] < b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
		log.Println(ret)
	} else {
		for i, j = 0, 0; i < len(a) && j < len(b); {
			if a[i] > b[j] {
				ret = append(ret, a[i])
				i++
			} else {
				ret = append(ret, b[j])
				j++
			}
		}
		log.Println(ret)
	}
	ret = append(ret, a[i:]...)
	ret = append(ret, b[j:]...)
	return ret
}

// the order of s[k] and v decide the sorted order of the array
func QuickSort(s []int, flag bool) []int {
	if len(s) <= 1 {
		return []int{}
	}
	var i, j, v, k = 0, len(s) - 1, s[0], 1

	for i < j {
		if flag { //big->small
			if s[k] < v {
				s[k], s[j] = s[j], s[k]
				j--
			} else {
				s[i], s[k] = s[k], s[i]
				i++
				k++
			}
		} else { //small -> big
			if s[k] > v {
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
	return s
}
