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
	up := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))
	sieve := make([]bool, (up+10)/2)
	for i := 3; i <= up; i += 2 {
		for j := i * i; j <= up; j += 2 * i {
			sieve[(j-1)/2] = true
		}
	}
	cnt := 0
	for i, v := range sieve {
		if !v {
			cnt++
			if n == cnt {
				return i*2 + 1
			}
		}
	}
	panic("unreachable")
	return 0
}
