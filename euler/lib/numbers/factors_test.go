package numbers

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