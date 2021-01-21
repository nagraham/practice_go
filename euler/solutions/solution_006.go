package euler

import "euler/lib/integers"

// Find the difference between the sum of the squares of the first
// one hundred natural numbers and the square of the sum.
func SumSquareDifference() int {
	start, end := 1, 100
	firstHundredNums := integers.Range(start, end+1)

	sumOfSquares := integers.Sum(integers.Map(firstHundredNums, func(n int) int {
		return n * n
	}))

	sum := integers.Sum(firstHundredNums)
	squareOfSum := sum * sum

	return squareOfSum - sumOfSquares
}
