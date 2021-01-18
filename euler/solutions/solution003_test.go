package euler

import (
	"fmt"
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	result := LargestPrimeFactor()
	fmt.Printf("The largest prime factor of 600851475143 = %d\n", result)
}