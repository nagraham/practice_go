package euler

import (
	"fmt"
	"testing"
)

func TestEvenFibonacciSum(t *testing.T) {
	sum := EvenFibonacciSum()
	fmt.Printf("Solution 2 answer: %d", sum)
}