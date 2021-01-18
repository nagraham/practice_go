package euler

import (
	"fmt"
	"testing"
)

func TestLargestPalindromeProduct(t *testing.T) {
	result := LargestPalindromeProduct()
	fmt.Printf("The largest palindrome product = %d\n", result)
}