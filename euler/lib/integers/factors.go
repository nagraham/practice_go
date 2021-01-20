package numbers

import (
	"math"
)

// Returns all factors of the given number as an ordered list of integers.
// e.g. 12 -> [1, 2, 3, 4, 6, 12]
func Factors(num int) []int {
	if num < 1 {
		return nil
	}

	// when you divide n cleanly with some number a, you get a dividend b; both are factors of n;
	// to help with ordering, store the a's in the low list, and b's in the high list; combine later
	var lowFactors []int
	var highFactors []int

	max := int(math.Round(math.Sqrt(float64(num))))
	for i := 1; i <= max; i++ {
		if num % i != 0 {
			continue
		}

		lowFactors = append(lowFactors, i)
		if num / i != i {
			highFactors = append(highFactors, num / i)
		}
	}

	return append(lowFactors, Reverse(highFactors)...)
}

// Returns the least common multiple for thw two integers, which is the smallest
// number that can be divided cleanly by both values.
// If either of the numbers are zero, or less, the result is -1
func LeastCommonMultiple(a int, b int) int {
	if a < 1 || b < 1 {
		return -1
	}
	return a * b / GreatestCommonDivisor(a, b)
}

// Returns the greatest common divisor, which is the largest number that
// cleanly divides both values.
// If either of the numbers are zero, or less, the result is -1
func GreatestCommonDivisor(a int, b int) int {
	if a < 1 || b < 1 {
		return -1
	}

	dividend := MaxInt(a, b)
	divisor := MinInt(a, b)
	remainder := -1

	for {
		remainder = dividend % divisor
		if remainder == 0 {
			break
		}
		dividend = divisor
		divisor = remainder
	}

	return divisor
}