package sorts

// determine the mid=i+(j-i)/2, and the bound is j=mid, i=mid+1
func BinarySearch(s []int, target int) int {
	var i, j, mid = 0, len(s) - 1, 0
	for i < j {
		mid = i + (j-i)/2
		if s[mid] == target {
			return mid
		} else if s[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return -1
}
