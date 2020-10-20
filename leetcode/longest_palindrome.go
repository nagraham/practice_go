package leetcode

// Find the longest palindrome substring
// Strings up to 1000 runes long; only english lower/upper case letters and numbers
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	palindrome := ""

	runeIndex := createRuneIndex(s)

	for r, indices := range runeIndex {

		// skip if rune only occurs once
		if len(indices) < 2 {
			continue
		}

		// check palindrome for each "pairing" of where the rune occurs;
		// skip if the length between the two runes is less than the current palindrome
		for i := 0; i < len(indices); i++ {
			for j := len(indices)-1; i < j; j-- {
				start := indices[i]
				end := indices[j] + len(string(r)) // string length is based on bytes; we need bytes of rune
				if len(palindrome) >= end - start {
					continue
				}

				if isPalindrome(s[start:end]) {
					palindrome = s[start:end]
				}
			}
		}
	}

	if palindrome == "" {
		palindrome = s[0:1]
	}

	return palindrome
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func createRuneIndex(s string) map[rune][]int {
	runeIndex := make(map[rune][]int)
	for i, r := range s {
		indices, ok := runeIndex[r]
		if !ok {
			indices = make([]int, 0)
		}
		runeIndex[r] = append(indices, i)
	}
	return runeIndex
}