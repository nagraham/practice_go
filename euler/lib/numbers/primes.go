package numbers

import "math"

var basePrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199,
}

// Quick test to see if a number is divisible by one of the first few primes
func isDivisibleBySmallPrime(num int, max int) bool {
	for _, p := range basePrimes {
		if p > max {
			break
		}
		if num % p == 0 {
			return true
		}
	}
	return false
}

// Part of an algorithm to jump through primes (see isDivisibleByLargerPrimeValues)
const initialK = 34

// this method is taken from this wiki:
// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func isDivisibleByLargerPrimeValues(num int, max int) bool {
	for k := initialK; k < max; k++ {
		k1, k2 := k*6-1, k*6+1
		if num % k1 == 0 || num % k2 == 0 {
			return true
		}
	}
	return false
}

// Returns true if the given number is prime.
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	max := int(math.Sqrt(float64(num)))

	// use short circuiting and negation to get our result
	return !(isDivisibleBySmallPrime(num, max) || isDivisibleByLargerPrimeValues(num, max))
}