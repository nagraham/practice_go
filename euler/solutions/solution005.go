package euler

import "euler/lib/numbers"

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func SmallestMultiple() int {

	// https://en.wikipedia.org/wiki/Least_common_multiple
	// the LCM method is associative: we can iterative through 1...20 to build up the LCM value
	lcm := 1
	for i := 2; i <= 20; i++ {
	  lcm = numbers.LeastCommonMultiple(lcm, i)
	}
	return lcm
}
