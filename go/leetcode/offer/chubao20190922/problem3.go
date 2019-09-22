package main

import (
	"fmt"
	"sort"
)

func problem3(n int, a []int) int {
	if n <= 0 || a == nil || len(a) == 0 {
		return 0
	}
	return len(threeSum(a))
}

func threeSum(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func problemWrapper3() {
	A, B := input3()
	output := problem3(A, B)
	fmt.Println(output)
}

func input3() (int, []int) {
	n := 0
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	return n, a
}

func main() {
	problemWrapper3()
}
