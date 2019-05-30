package sorts

import (
	"log"
)

// i indicates the sort iteration, j indicates the swap operation
func BubleSort(s []int, flag bool) []int {
	for i, _ := range s {
		for j := 0; j < len(s)-i-1; j++ {
			if flag && s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			} else if !flag && s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
		log.Println(s)
	}
	log.Println("Sort Over!")
	return s
}

// iSorted indicates the last sorted index
// introduce a temp to store the cache value to insert or use [iSorted+1]
// in the sorted part, use binarySearch to find the postion is quicker than iterate to find
func InsertSort(s []int, flag bool) []int {
	var temp int
	for iSorted := 0; iSorted < len(s)-1; iSorted++ {
		if !flag && s[iSorted] > s[iSorted+1] {
			temp = s[iSorted+1]
			for j := iSorted; j >= 0; j-- {
				if s[j] > temp {
					s[j+1] = s[j]
					s[j] = temp
				} else {
					s[j+1] = temp
					break
				}
			}
		} else if flag && s[iSorted] < s[iSorted+1] {
			temp = s[iSorted+1]
			for j := iSorted; j >= 0; j-- {
				if s[j] < temp {
					s[j+1] = s[j]
					s[j] = temp
				} else {
					s[j+1] = temp
				}
			}
		}
		log.Println(s)
	}
	return s
}

func InsertReview(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] { //small->big
			temp := s[i+1]
			for j := i; j >= 0; j-- {
				if s[j] < temp {
					s[j+1] = temp
					break
				} else {
					s[j], s[j+1] = s[j+1], s[j]
					if j == 0 {
						s[j] = temp
					}
				}
			}
		}
	}
	return s
}

// use temp to instore the temp position of the min/max value of the unsorted part
// always select the min/max value from the unsorted part
func SelectSort(s []int, flag bool) []int {
	var temp int
	for i := 0; i < len(s)-1; i++ {
		temp = i
		if flag {
			for j := i + 1; j < len(s); j++ {
				if s[j] < s[temp] {
					temp = j
				}
			}
			s[i], s[temp] = s[temp], s[i]
		} else {
			for j := i + 1; j < len(s); j++ {
				if s[j] > s[temp] {
					temp = j
				}
			}
			s[i], s[temp] = s[temp], s[i]
		}
		log.Println(s)
	}
	return s
}

//
func SelectReview(s []int) []int {
	var k int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j] < s[i] {
				k = j
			}
		}
		s[i], s[k] = s[k], s[i]
	}
	return s
}
