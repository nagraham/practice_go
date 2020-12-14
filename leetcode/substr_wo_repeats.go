package leetcode

import (
	"math"
)

/*

type CharState struct: {
  start int		// index of character
  end int		// index of repeated character; default to -1
}

map[char]CharState

for each character in string
  if char has CharState
    if cs.end == -1
       this is the first repeat, set index
    els
 */

/*
Brute force:
nested loops, iterate each possible substring
Keep set of characters to track if you have encountered letter
when you encounter duplicate letter, terminate loop.
Check the substring length against the length of max substring discovered so far. Replace if longer.
Return the max substring

You could use an array instead of a set ... b
ut the instructions don't specify a char set.
Just english letters, spaces, symbols.
*/


/*
Decompose:

(A) Find max length substring
  - Cache longest substring discovered
  - Reduce against the length of each candidate
  - CompareCache ... a general solution ... cache single entity, use given compare function to compare w/ cache

func reduceFunction should have signature:  (a interface{}, b interface{}) interface{} ... such that it returns a or b

NewCompareAndCache(func reduceFunction)	// Creates new CompareAndCache w/ given reduceFunction
Reduce(x interface{})					// Compares given value against the cached value
Get() interface{}						// Gets the cached value


(B) Iterate and find substring without duplicates
  - Each character in string
  - Add to set to track discovered characters
  - If char in set, we found a duplicate
  - Optimize:
    - Set also tracks indices of characters (so it is a map[char]int)
    - Track start/end indices of current substring
    - When you find a duplicate, compare substring against max, etc, then ...
      update char's index with the newly discovered index. You effectively replace the instance
      of character with the newly discovered one.
    - But before replacing, set the "start" index pointer to the old duplicate's index + 1
    - This gets you O(n) instead of O(n^2 / 2)
 */

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	runeMap := make(map[rune]int)

	start := 0
	for i, r := range s {
		lastOccurrence, hasDuplicate := runeMap[r]

		// if lastOccurrence < start, we have already found another duplicate rune
		if hasDuplicate && lastOccurrence >= start {
			maxLen = MaxInt(maxLen, len(s[start:i]))
			start = lastOccurrence + 1
		}
		runeMap[r] = i
	}

	maxLen = MaxInt(maxLen, len(s[start:]))

	return maxLen
}

func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}