package main

import (
	// c "./contestSet"
	// p "./problemSet"
	"log"
)

func reverseString(s string) string {
	var b []byte = []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[j], b[i] = b[i], b[j]
	}
	return string(b)
}

// using map to solve
func twoSum1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		if value, ok := m[target-v]; ok {
			return []int{value + 1, i + 1}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

// two pointer solution
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{left + 1, right + 1}
}

func removeElement(nums []int, val int) int {
	j := 0
	for _, v := range nums {
		if v != val {
			nums[j] = v
			j++
		}
	}
	nums = nums[:j]
	log.Println(nums)
	return len(nums)
}

func findMaxConsecutiveOnes(nums []int) int {
	nums = append([]int{0}, nums...)
	nums = append(nums, 0)
	max, j := 0, 0
	for i, v := range nums {
		if v == 0 && i-j > max {
			max = i - j - 1
		}
		if v == 0 {
			j = i
		}
	}
	return max
}

func minSubArrayLen(s int, nums []int) int {
	sum, i, j, min := 0, 0, 0, len(nums)+1
	if min == 1 {
		return 0
	}
	for _, v := range nums {
		sum += v
		j++
		for sum >= s {
			if j-i < min {
				min = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	log.Println(nums)
}

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	} else {
		rowIndex++
	}
	front, res := []int{1}, []int{}
	for i := 1; i < rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res = append(res, 1)
			} else {
				res = append(res, front[j-1]+front[j])
			}
		}
		front = res
		res = []int{}
	}
	return front
}

func main() {
	// res := p.AddBinary("110010", "10111")
	// res := p.StrStr("hello", "ll")
	// res := reverseString("hello")
	// res := twoSum([]int{2, 7, 11, 15}, 9)
	// res := twoSum([]int{2, 3, 4}, 6)
	// res := twoSum([]int{2, 3, 4, 5, 11, 17, 19}, 16)
	// res := twoSum([]int{5, 25, 75}, 100)
	// res := removeElement([]int{3, 2, 2, 3, 5, 6, 2, 5}, 2)
	// res := removeElement([]int{3, 2, 2, 3, 5, 6, 2, 5}, 1)
	// res := removeElement([]int{}, 2)
	// res := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0})
	// res := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	// res := findMaxConsecutiveOnes([]int{0, 0, 0, 0, 1, 1, 0, 1, 1, 1})
	// res := findMaxConsecutiveOnes([]int{0, 0, 0, 0})
	// res := findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1})
	// res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	// rotate([]int{2, 3, 1, 2, 4, 3}, 1)
	// rotate([]int{1, 2, 3, 4, 5, 6, 7}, 10)
	// res := getRow(0)
	// res := getRow(1)
	// log.Printf("%d\n", res)
	// res = getRow(2)
	// log.Printf("%d\n", res)
	// res = getRow(3)
	// log.Printf("%d\n", res)
	// res = getRow(4)
	res := getRow(33)
	log.Printf("%d\n", res)
}
