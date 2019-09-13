package main

import (
	"fmt"
	"math"
	"strconv"
)

// 判断是否为素数
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	k := int(math.Sqrt(float64(n))) + 1
	for i := 5; i < k; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 判断输入是否合法
func isLegal(s string) bool {
	if len(s) != 6 {
		return false
	}
	for i := range s {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return false
		}
	}
	n, err := strconv.Atoi(s)
	// fmt.Println(n)
	if err != nil {
		return false
	}
	return isPrime(n)
}

// 判断一次转换是否有效
func isTransable(s string, i int) (bool, int) {
	bs := []byte(s)
	for j := 0; j < 10; j++ {
		bs[i] = byte('0' + j)
		if num, err := strconv.Atoi(string(bs)); err != nil {
			break
		} else if isPrime(num) {
			return true, j
		}
	}
	return false, -1
}

func problem2(A, B []string, n int) []int {
	res := make([]int, n)
	// 处理每一个字符
	for i := 0; i < n; i++ {
		// 判断合法
		if !isLegal(A[i]) || !isLegal(B[i]) {
			res[i] = -1
			continue
		}
		// 找不同的位置
		diff := make([]int, 0, 6)
		for k := 0; k < 6; k++ {
			if A[i][k] != B[i][k] {
				diff = append(diff, k)

			}
		}
		// 替换过滤，回溯法替换过滤？
		temp := 0
		for _, k := range diff {
			// 按照位置进行转换
			if flag, j := isTransable(A[i], k); flag {
				temp++
				A[i] = A[i][:k] + string(byte('0'+j)) + A[i][k+1:]
			} else {
				temp = -1
				break
			}
		}
		res[i] = temp
	}
	return res
}

func problemWrapper2() {
	A, B, n := input2()
	output := problem2(A, B, n)
	fmt.Println(output)
	// for _, out := range output {
	// 	fmt.Println(out)
	// }
}

func input2() ([]string, []string, int) {
	n := 0
	fmt.Scanf("%d", &n)
	A, B := make([]string, n), make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &A[i], &B[i])
	}
	return A, B, n
}

func main() {
	problemWrapper2()
}
