package euler

import (
	"euler/lib/numbers"
	"euler/lib/strings"
	"strconv"
)

const maxNum = 999

// Find the largest palindrome made from the product of two 3-digit numbers.
func LargestPalindromeProduct() int {
	largestPalindrome := -1

	// This nested loop works through columns of the multiplication table,
	// from highest to lowest values; it allows us to avoid duplicate
	// products (999 * 998 & 998 * 999), and we work on smaller "chunks" of products,
	// with the best candidates for our result. We don't need to cover all 999^2 products.
	for i := 0; i < maxNum; i++ {
		// once we encounter a column whose largest number is less than the best palindrome we've found,
		// no remaining column will contain a greater palindrome, and we can complete.
		if (maxNum-i) * maxNum < largestPalindrome {
			return largestPalindrome
		}
		for j := 0; j <= i; j++ {
			prod := (maxNum-i) * (maxNum-j)
			if strings.IsPalindrome(strconv.Itoa(prod)) {
				largestPalindrome = numbers.MaxInt(largestPalindrome, prod)
			}
		}
	}

	return -1
}

