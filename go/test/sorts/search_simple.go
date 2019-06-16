package sorts

func SeqSearch(a []int, target int) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}

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

//
func FibSearch(a []int, target int) int {
	var n int = len(a)
	// Auto generate Fibonacci seq & Find the first k meet n<F[k]-1
	F, k := func(n int) ([]int, int) {
		var dp []int = []int{}
		dp = append(dp, 1) //dp[0]=1
		dp = append(dp, 1) //dp[1]=1
		var i int = 2
		for ; dp[i-1]-1 < n; i++ {
			dp = append(dp, dp[i-1]+dp[i-2]) //dp[i]=dp[i-1]+dp[i-2]
		}
		return dp, i - 1
	}(n)
	// Extend the original a to meet up F[k]-1
	for i := n; i < F[k]-1; i++ {
		a = append(a, a[i-1])
	}
	// do the Search
	var i, j, mid int = 0, len(a) - 1, 0
	for i <= j {
		mid = i + F[k-1] - 1
		switch {
		case target == a[mid]:
			if mid < n {
				return mid
			} else {
				return n - 1
			}
		case target < a[mid]:
			j = mid - 1
			k -= 1
		default: //target > a[mid]
			i = mid + 1
			k -= 2
		}
	}
	return -1
}

func HashSearch(a []int, target int) int {
	m := make(map[int][]int, len(a))
	// preprocess
	for i, v := range a {
		m[v] = append(m[v], i)
	}
	// search the table
	i, ok := m[target]
	if ok {
		return i[0]
	}
	return -1
}
