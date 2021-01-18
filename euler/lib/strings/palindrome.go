package strings

func IsPalindrome(s string) bool {
	for start, end := 0, len(s)-1; start <= end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}

	return true
}
