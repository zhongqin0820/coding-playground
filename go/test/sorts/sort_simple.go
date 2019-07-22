package sorts

// 边界判断代码
func Defensive(a []int) bool {
	if a == nil || len(a) == 0 || len(a) == 1 {
		return true
	}
	return false
}

// buble sort
// i indicates the sort iteration, j indicates the swap operation
func BubleSort(s []int, flag bool) {
	if Defensive(s) {
		return
	}
	for i, _ := range s {
		for j := 0; j < len(s)-i-1; j++ {
			if flag && s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			} else if !flag && s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// Select sort
// use temp to instore the temp position of the min/max value of the unsorted part
// always select the min/max value from the unsorted part
func SelectSort(s []int, flag bool) {
	if Defensive(s) {
		return
	}
	var temp int
	for i := 0; i < len(s)-1; i++ {
		temp = i
		if !flag {
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
	}
}

// insert sort
func InsertSortNavie(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	n := len(a)
	for i := 1; i < n; i++ {
		if !flag && a[i] < a[i-1] { // small -> big
			for j := i; j > 0 && a[j] < a[j-1]; j-- {
				a[j], a[j-1] = a[j-1], a[j]
			}
		} else if flag && a[i] > a[i-1] { // big -> small
			for j := i; j > 0 && a[j] > a[j-1]; j-- {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// iSorted indicates the last sorted index
// introduce a temp to store the cache value to insert or use [iSorted+1]
// in the sorted part, use binarySearch to find the postion is quicker than iterate to find
func InsertSort(s []int, flag bool) {
	if Defensive(s) {
		return
	}
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
					break
				}
			}
		}
	}
}

// shell sort is the upgrade version of insert sort
func ShellSort(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	n := len(a)
	// split the original array into gap groups
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			if !flag && a[i] < a[i-gap] { // small -> big
				for j := i; j-gap >= 0 && a[j] < a[j-gap]; j -= gap {
					a[j], a[j-gap] = a[j-gap], a[j]
				}
			} else if flag && a[i] > a[i-gap] { // big-> small
				for j := i; j-gap >= 0 && a[j] > a[j-gap]; j -= gap {
					a[j], a[j-gap] = a[j-gap], a[j]
				}
			}
		}
	}
}
