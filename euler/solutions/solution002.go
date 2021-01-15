package euler

import (
	"euler/lib/fibonacci"
	"euler/lib/nums"
)

func EvenFibonacciSum() int {
	// get sequence of fibonacci up to 4000000
	fibs := fibonacci.UpTo(4000000)

	// filter out odd numbers
	evenFibs := nums.Filter(fibs, func(num int) bool { return num % 2 == 0 })

	// sum the numbers
	return nums.Sum(evenFibs)
}
