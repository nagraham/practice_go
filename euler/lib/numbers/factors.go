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