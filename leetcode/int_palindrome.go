package leetcode

import "strconv"

/*
Determine if an integer is a palindrome.
Negative number is not a palindrome: -121 -> 121-

Challenge:
Could you solve it without converting to a string?
- Yes, mod 10, add number to slice, divide by 10, until 0, then use palindrome algorithm on array
- But this wouldn't be as simple to solve
*/

func isIntPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return isPalindrome(strconv.Itoa(x))
}