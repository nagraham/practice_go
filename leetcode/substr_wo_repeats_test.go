package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring_NoDuplicates(t *testing.T) {
	result := lengthOfLongestSubstring("abcdefghij")
	assert.Equal(t, 10, result)
}

func TestLengthOfLongestSubstring_DuplicateAtEnd(t *testing.T) {
	result := lengthOfLongestSubstring("abcdefghija")
	assert.Equal(t, 10, result)
}

func TestLengthOfLongestSubstring_MultipleDuplicates(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, result)
}

func TestLengthOfLongestSubstring_AllDuplicates(t *testing.T) {
	result := lengthOfLongestSubstring("bbbbb")
	assert.Equal(t, 1, result)
}

func TestLengthOfLongestSubstring_ConsecutiveDuplicates(t *testing.T) {
	result := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, result)
}

func TestLengthOfLongestSubstring_TriggerSubstringCheckDifferentRune(t *testing.T) {
	result := lengthOfLongestSubstring("abcdefghijc")
	assert.Equal(t, 10, result)
}

func TestLengthOfLongestSubstring_EmptyString(t *testing.T) {
	result := lengthOfLongestSubstring("")
	assert.Equal(t, 0, result)
}

func TestLengthOfLongestSubstring_LongestSubstringAtEndHalf(t *testing.T) {
	result := lengthOfLongestSubstring("aab")
	assert.Equal(t, 2, result)
}

func TestLengthOfLongestSubstring_EncounterEarlierDuplicate(t *testing.T) {
	result := lengthOfLongestSubstring("abba")
	assert.Equal(t, 2, result)
}