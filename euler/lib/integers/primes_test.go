package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	testConditions := map[string]struct {
		num int
		expectedResult bool
	} {
		"when zero": {
			num: 0,
			expectedResult: false,
		},
		"when one": {
			num: 1,
			expectedResult: false,
		},
		"when two": {
			num: 2,
			expectedResult: true,
		},
		"when three": {
			num: 3,
			expectedResult: true,
		},
		"when a small non-prime": {
			num:            4,
			expectedResult: false,
		},
		"when a medium non-prime": {
			num: 92,
			expectedResult: false,
		},
		"when a medium prime": {
			num: 97,
			expectedResult: true,
		},
		"when a large non-prime": {
			num: 450,
			expectedResult: false,
		},
		"when a large prime": {
			num: 331,
			expectedResult: true,
		},
		"when a very large non-prime": {
			num: 15000003,
			expectedResult: false,
		},
		"when a very large prime": {
			num: 300007,
			expectedResult: true,
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := IsPrime(cond.num)
			assert.Equal(t, cond.expectedResult, result)
		})
	}
}

func TestNthPrime(t *testing.T) {
	testConditions := map[string]struct {
		n int
		expectedPrime int
	} {
		"when zero": {
			n: 0,
			expectedPrime: 0,
		},
		"when one": {
			n: 1,
			expectedPrime: 2,
		},
		"when two": {
			n: 2,
			expectedPrime: 3,
		},
		"when fifty": {
			n: 50,
			expectedPrime: 229,
		},
		"when fifty-one": {
			n: 51,
			expectedPrime: 233,
		},
		"when seventy-seven": {
			n: 77,
			expectedPrime: 389,
		},
		"when one-hundred": {
			n: 100,
			expectedPrime: 541,
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := NthPrime(cond.n)
			assert.Equal(t, cond.expectedPrime, result)
		})
	}
}