package web

import (
	"fmt"
	"testing"
)

// 函数类型
type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 != 0 {
		return false
	}
	return true
}

// 函数式编程: 回调
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

//
func Testfilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice= ", slice)
	// 将函数作为参数传入
	odd := filter(slice, isOdd)
	fmt.Println("odd slice= ", slice)
	//
	even := filter(slice, isEven)
	fmt.Println("even slice= ", slice)
}
