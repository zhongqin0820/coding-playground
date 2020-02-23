package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/prime-arrangements/
var helper1175 = 1000000007

func numPrimeArrangements(n int) int {
	prime, nonPrime, res := 0, 0, 1
	// prime sieve
	var sieve = make([]bool, n+1)
	sieve[0], sieve[1] = true, true
	for i := 2; i*i <= n; i++ {
		if !sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	for i := range sieve {
		if !sieve[i] {
			prime++
		}
	}
	nonPrime = n - prime
	// construct result
	for i := 2; i <= prime; i++ {
		res *= i
		res %= helper1175
	}
	for i := 2; i <= nonPrime; i++ {
		res *= i
		res %= helper1175
	}
	return res
}

func TestNumPrimeArrangements(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{5, 12},
		{100, 682289015},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, numPrimeArrangements(ts.n))
		})
	}
}
