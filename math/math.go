package math

import "math"

var smallPrimes = []int{0, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func NthPrime(n int) int {

	if n < 1 {
		panic("error")
	}

	if n < len(smallPrimes) {
		return smallPrimes[n]
	}

	up := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))

	sieve := make([]bool, up/2)

	cnt := 1
	for i := 3; i < up; i += 2 {
		if sieve[(i-1)/2] == false {
			cnt++
			if n == cnt {
				return i
			}
			for j := i * i; j < up; j += 2 * i {
				sieve[(j-1)/2] = true
			}
		}
	}

	panic("unreachable")
	return 0
}
