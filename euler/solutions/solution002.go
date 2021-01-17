package euler

import (
	"euler/lib/fibonacci"
	"euler/lib/numbers"
)

func EvenFibonacciSum() int {
	// get sequence of fibonacci up to 4000000
	fibs := fibonacci.UpTo(4000000)

	// filter out odd numbers
	evenFibs := numbers.Filter(fibs, numbers.IsEven)

	// sum the numbers
	return numbers.Sum(evenFibs)
}

func EvenFibonacciSumTwo() int {
	var sum int
	a := 0
	b := 1
	for b < 4000000 {
		if b % 2 == 0 {
			sum = sum + b
		}
		tmp := b
		b = b + a
		a = tmp
	}
	return sum
}

func EvenFibonacciSumThree() int {
	fibs := fibonacciSequence(4000000)
	return sumIfEven(fibs)
}


func fibonacciSequence(max int) []int {
	if max == 0 {
		return []int{}
	}
	if max == 1 {
		return []int{0}
	}

	fibs := []int{0, 1}

	for next(fibs) < max {
		fibs = append(fibs, next(fibs))
	}

	return fibs
}

func next(fibs []int) int {
	return fibs[len(fibs) - 1] + fibs[len(fibs) - 2]
}

func sumIfEven(nums []int) int {
	var sum int
	for _, num := range nums {
		if num % 2 == 0 {
			sum += num
		}
	}
	return sum
}