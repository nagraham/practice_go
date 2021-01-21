package euler

import (
	"fmt"
	"testing"
)

func TestEvenFibonacciSum(t *testing.T) {
	sum := EvenFibonacciSum()
	fmt.Printf("Solution 2 answer: %d\n", sum)
}

func TestEvenFibonacciSumTwo(t *testing.T) {
	sum := EvenFibonacciSumTwo()
	fmt.Printf("Solution 2.2 answer: %d\n", sum)
}

func TestEvenFibonacciSumThree(t *testing.T) {
	sum := EvenFibonacciSumThree()
	fmt.Printf("Solution 2.3 answer: %d\n", sum)
}