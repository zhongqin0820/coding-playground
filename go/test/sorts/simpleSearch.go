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

func InterpolationSearch(a []int, target int) int {
	i, mid, j, n := 0, 0, len(a)-1, len(a)-1
	for i <= j {
		mid = i + (j-i)*(target-a[i])/(a[j]-a[i])
		if mid < 0 || mid > n {
			break
		}
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
