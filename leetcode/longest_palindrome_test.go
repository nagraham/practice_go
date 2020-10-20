package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome_Palindromes(t *testing.T) {
	assert.True(t, isPalindrome("aba"))
	assert.True(t, isPalindrome("racecar"))
	assert.True(t, isPalindrome("hannah"))
}

func TestIsPalindrome_NotPalindromes(t *testing.T) {
	assert.False(t, isPalindrome("abb"))
	assert.False(t, isPalindrome("carrace"))
	assert.False(t, isPalindrome("nah nah"))
}

func TestLongestPalindrome_emptyString(t *testing.T) {
	result := longestPalindrome("")
	assert.Equal(t, "", result)
}

func TestLongestPalindrome_singleRune(t *testing.T) {
	result := longestPalindrome("a")
	assert.Equal(t, "a", result)
}

func TestLongestPalindrome_LengthTwo_DifferentRunes(t *testing.T) {
	result := longestPalindrome("ab")
	assert.Equal(t, "a", result)
}

func TestLongestPalindrome_LengthTwo_SameRunes(t *testing.T) {
	result := longestPalindrome("aa")
	assert.Equal(t, "aa", result)
}

func TestLongestPalindrome_PalindromeWithinShortString(t *testing.T) {
	result := longestPalindrome("babad")
	assertIsOneOf(t, result, "aba", "bab")
}

func TestLongestPalindrome_PalindromeInMiddleOfShortString(t *testing.T) {
	result := longestPalindrome("cbbd")
	assert.Equal(t, "bb", result)
}

func TestLongestPalindrome_WholeStringIsPalindrome(t *testing.T) {
	result := longestPalindrome("racecar")
	assert.Equal(t, "racecar", result)
}

func TestLongestPalindrome_MultiplePalindromesDifferentLengths(t *testing.T) {
	result := longestPalindrome("abcbabcddcba")
	assert.Equal(t, "abcddcba", result)
}

/*
 *  HELPER FUNCTIONS
 */

func assertIsOneOf(t *testing.T, actual string, expected... string) {
	assert.Subset(t, expected, []string{ actual })
}
