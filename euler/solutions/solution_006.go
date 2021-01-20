package euler

import "euler/lib/numbers"

// Find the difference between the sum of the squares of the first
// one hundred natural numbers and the square of the sum.
func SumSquareDifference() int {
	start, end := 1, 100
	firstHundredNums := numbers.Range(start, end+1)

	sumOfSquares := numbers.Sum(numbers.Map(firstHundredNums, func(n int) int {
		return n * n
	}))

	sum := numbers.Sum(firstHundredNums)
	squareOfSum := sum * sum

	return squareOfSum - sumOfSquares
}
