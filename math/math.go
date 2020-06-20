package math

import "math"

var smallPrimes = []int{0, 2, 3, 5, 7, 11, 13}

func NthPrime(n int) int {

	if n < 1 {
		panic("error")
	}

	if n < len(smallPrimes) {
		return smallPrimes[n]
	}

	up := int(float64(n)*(math.Log(float64(n))+math.Log(float64(n)))) + 2

	sieve := make([]bool, up+10)

	for i := 2; i < up; i++ {
		for j := i * i; j < up; j += i {
			sieve[j] = true
		}
	}

	cnt := 0
	for i, v := range sieve {
		if !v {
			cnt++
		}
		if n == cnt-2 {
			return i
		}
	}

	panic("unreachable")
	return 0
}
