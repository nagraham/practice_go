package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIntPalindrome_Zero(t *testing.T) {
	result := isIntPalindrome(0)
	assert.True(t, result)
}

func TestIntPalindrome_NegativeNumbersAreAllFalse(t *testing.T) {
	for _, num := range []int{ -1, -44, -121, -1221, math.MinInt32 } {
		assert.False(t, isIntPalindrome(num))
	}
}

func TestIntPalindrome_palindromeNumbers(t *testing.T) {
	for _, num := range []int{ 1, 121, 1221, 333, 1234554321 } {
		assert.True(t, isIntPalindrome(num))
	}
}