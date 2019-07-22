package sorts

import (
	"sort"
)

func CountingSort(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	// step 1: find the min & max of a
	min, max, n := a[0], a[0], len(a)
	// step 2: generates a map range [min, max]
	m := map[int]int{a[0]: 1}
	//
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		m[a[i]]++
	}
	// step 3: reassign
	if !flag { // small -> big
		for i, j := min, 0; i <= max; i++ {
			if lenk, ok := m[i]; ok {
				for k := 0; k < lenk; k++ {
					a[j] = i
					j++
				}
			}
		}
	} else { // big->small
		for i, j := max, 0; i >= min; i-- {
			if lenk, ok := m[i]; ok {
				for k := 0; k < lenk; k++ {
					a[j] = i
					j++
				}
			}
		}
	}
}

func BucketSort(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	n, m := len(a), map[int][]int{}
	// step 1: find the min & max of a
	min, max := a[0], a[0]
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	buckets := max - min
	// step 2: maps the a[i] into different buckets
	for i := 0; i < n; i++ {
		m[(a[i]/buckets)%buckets] = append(m[(a[i]/buckets)%buckets], a[i])
	}
	// step 3: collects the buffered buckets
	for k := 0; k < buckets; k++ {
		if _, ok := m[k]; ok {
			// step 4: sorted each buckets
			if !flag { // small -> big
				sort.Slice(m[k], func(i, j int) bool {
					return m[k][i] < m[k][j]
				})
			} else { // big -> small
				sort.Slice(m[k], func(i, j int) bool {
					return m[k][i] > m[k][j]
				})
			}
		}
	}
	// step 5: merge each bucket
	if !flag {
		for i, j := 0, 0; i <= buckets; i++ {
			if _, ok := m[i]; ok {
				for k := 0; k < len(m[i]); k, j = k+1, j+1 {
					a[j] = m[i][k]
				}
			}
		}
	} else {
		for i, j := buckets, 0; i >= 0; i-- {
			if _, ok := m[i]; ok {
				for k := 0; k < len(m[i]); k, j = k+1, j+1 {
					a[j] = m[i][k]
				}
			}
		}
	}
}

func RadixSort(a []int, flag bool) {
	if Defensive(a) {
		return
	}
	// step 1: find max & its highest bits
	max, bits, n := a[0], 0, len(a)
	for i := 1; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	for temp := max; temp != 0; temp /= 10 {
		bits++
	}
	//
	for i, j, k, radix := 0, 0, 0, 1; i < bits; i++ {
		buckets := new([10]int)
		tmp := make([]int, n, n)
		// step 2: map into buckets
		for j = 0; j < n; j++ {
			k = (a[j] / radix) % 10
			buckets[k]++
		}
		// step 3: counting sort
		for j = 1; j < 10; j++ {
			buckets[j] = buckets[j-1] + buckets[j]
		}
		for j = n - 1; j >= 0; j-- {
			k = (a[j] / radix) % 10
			tmp[buckets[k]-1] = a[j]
			buckets[k]--
		}
		// step 4: reassign
		for j = 0; j < n; j++ {
			a[j] = tmp[j]
		}
		radix = radix * 10
	}
	// big -> small
	if flag {
		for i := 0; i < n/2; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
	}
}
