package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/prime-palindrome/

// Find the smallest prime palindrome greater than or equal to N.
// Recall that a number is prime if it's only divisors are 1 and itself, and it is greater than 1.
// For example, 2,3,5,7,11 and 13 are primes.
// Recall that a number is a palindrome if it reads the same from left to right as it does from right to left.
// For example, 12321 is a palindrome.

// Note:
// 1 <= N <= 10^8
// The answer is guaranteed to exist and be less than 2 * 10^8.

func primePalindrome(N int) int {
	for {
		if helper866isPrime(N) && helper866isPalindrome(N) {
			return N
		}
		N++
		if N > 1000 && N < 10000 { //10e4
			N = 10000
		}
		if N > 100000 && N < 1000000 { //10e6
			N = 1000000
		}
		if N > 10000000 && N < 100000000 { //10e8
			N = 100000000
		}

	}
	return -1
}

// 判断一个数字是否是回文
func helper866isPalindrome(n int) bool {
	temp, res := n, 0
	for temp > 0 {
		res = res*10 + temp%10
		temp /= 10
	}
	return n == res
}

//判断一个数字是否是素数
func helper866isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestPrimePalindrome(t *testing.T) {
	tests := []struct {
		N      int
		output int
	}{
		{6, 7},
		{8, 11},
		{13, 101},
		{1000, 10301},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, primePalindrome(ts.N))
		})
	}
}
