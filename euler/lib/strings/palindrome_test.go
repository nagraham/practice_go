package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testConditions := map[string]struct {
		s string
		expectedResult bool
	} {
		"empty string": {
			s: "",
			expectedResult: true,
		},
		"one letter": {
			s: "a",
			expectedResult: true,
		},
		"all whitespace": {
			s: "   ",
			expectedResult: true,
		},
		"palindrome": {
			s: "racecar",
			expectedResult: true,
		},
		"non-plaindrome": {
			s: "racecars",
			expectedResult: false,
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := IsPalindrome(cond.s)
			assert.Equal(t, cond.expectedResult, result)
		})
	}
}