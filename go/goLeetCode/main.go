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

func main() {
	// res := p.AddBinary("110010", "10111")
	// res := p.StrStr("hello", "ll")
	// res := reverseString("hello")
	// res := twoSum([]int{2, 7, 11, 15}, 9)
	// res := twoSum([]int{2, 3, 4}, 6)
	// res := twoSum([]int{2, 3, 4, 5, 11, 17, 19}, 16)
	// res := twoSum([]int{5, 25, 75}, 100)
	// res := removeElement([]int{3, 2, 2, 3, 5, 6, 2, 5}, 2)
	res := removeElement([]int{3, 2, 2, 3, 5, 6, 2, 5}, 1)
	// res := removeElement([]int{}, 2)
	log.Printf("%d\n", res)
}
