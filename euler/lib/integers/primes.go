package integers

import "math"

var basePrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
}

// Part of an algorithm to jump through primes (see isDivisibleByLargerPrimeValues)
const initialK = 39

// Returns true if the given number is prime.
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	// use short circuiting and negation to get our result
	return isPrime(num, basePrimes) && !isDivisibleByLargerPrimeValues(num)
}

// Returns the Nth prime number, where 1 < n < ?
// If n < 1, returns 0
func NthPrime(n int) int {
	if n < 1 {
		return 0
	}
	if n-1 < len(basePrimes) {
		return basePrimes[n - 1]
	}

	var primes []int
	primes = append(primes, basePrimes...)

	// the "k" approach inspired by https://en.wikipedia.org/wiki/Primality_test#Simple_methods
	for k := initialK; len(primes) < n; k++ {
		k1, k2 := k*6-1, k*6+1

		if isPrime(k1, primes) {
			primes = append(primes, k1)
		}
		if isPrime(k2, primes) {
			primes = append(primes, k2)
		}
	}

	return primes[n-1]
}

// Basic primality test based on whether the number can be divided evenly
// by any of the primes in the given list, up to the sqrt(num)
func isPrime(num int, primes []int) bool {
	max := int(math.Sqrt(float64(num)))
	for _, p := range primes {
		if p > max {
			break
		}
		if num % p == 0 {
			return false
		}
	}
	return true
}

// Helper method with checking if larger numbers whose sqrt > 229.
// This method is taken from this wiki:
// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func isDivisibleByLargerPrimeValues(num int) bool {
	max := int(math.Sqrt(float64(num)))

	for k := initialK; k < max; k++ {
		k1, k2 := k*6-1, k*6+1
		// one, both, or neither might be prime; either way, if cleanly divded, the number is not prime.
		if num % k1 == 0 || num % k2 == 0 {
			return true
		}
	}
	return false
}