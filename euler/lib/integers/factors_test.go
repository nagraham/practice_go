package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactors(t *testing.T) {
	testConditions := map[string]struct{
		number int
		expectedFactors []int
	} {
		"with zero": {
			number: 0,
			expectedFactors: nil,
		},
		"with one": {
			number: 1,
			expectedFactors: []int{1},
		},
		"with a prime": {
			number: 11,
			expectedFactors: []int{1, 11},
		},
		"with four": {
			number: 4,
			expectedFactors: []int{1, 2, 4},
		},
		"with twenty-four": {
			number: 24,
			expectedFactors: []int{1, 2, 3, 4, 6, 8, 12, 24},
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {

			result := Factors(cond.number)
			assert.Equal(t, cond.expectedFactors, result)
		})
	}
}

func TestGreatestCommonDivisor(t *testing.T) {
	testConditions := map[string]struct {
		a int
		b int
		gcd int
	} {
		"with two primes": {
			a: 2,
			b: 3,
			gcd: 1,
		},
		"with two non-prime numbers, one large, one smaller": {
			a: 48,
			b: 18,
			gcd: 6,
		},
		"with two non-prime numbers, one smaller, one large": {
			a: 18,
			b: 48,
			gcd: 6,
		},
		"with a zero in 'a' param": {
			a: 0,
			b: 12,
			gcd: -1,
		},
		"with a zero in 'b' param": {
			a: 8,
			b: 0,
			gcd: -1,
		},
		"with a negative number in 'a' param": {
			a: -8,
			b: 12,
			gcd: -1,
		},
		"with a negative number in 'b' param": {
			a: 8,
			b: -12,
			gcd: -1,
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := GreatestCommonDivisor(cond.a, cond.b)
			assert.Equal(t, cond.gcd, result)
		})
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	testConditions := map[string]struct {
		nums []int
		lcm int
	} {
		"with two non-prime numbers": {
			nums: []int{21, 6},
			lcm: 42,
		},
		"test transitivity": {
			nums: []int{3, 4, 6},
			lcm: 12,
		},
		"with a zero in 'a' param": {
			nums: []int{0, 12},
			lcm: -1,
		},
		"with a zero in 'b' param": {
			nums: []int{12, 0},
			lcm: -1,
		},
		"with a negative number in 'a' param": {
			nums: []int{-1, 12},
			lcm: -1,
		},
		"with a negative number in 'b' param": {
			nums: []int{12, -1},
			lcm: -1,
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			if len(cond.nums) > 2 {
				resultOne := LeastCommonMultiple(cond.nums[0], cond.nums[1])
				resultTwo := LeastCommonMultiple(cond.nums[2], resultOne)
				assert.Equal(t, cond.lcm, resultTwo)
			} else {
				result := LeastCommonMultiple(cond.nums[0], cond.nums[1])
				assert.Equal(t, cond.lcm, result)
			}
		})
	}
}