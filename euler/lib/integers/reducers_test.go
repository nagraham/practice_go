package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	testConditions := map[string]struct {
		nums []int
		expectedSum int
	} {
		"empty slice": {
			nums: []int{},
			expectedSum: 0,
		},
		"one number": {
			nums: []int{ 5 },
			expectedSum: 5,
		},
		"a few numbers": {
			nums: []int{ 1, 2, 3, 4, 5},
			expectedSum: 120,
		},
	}

	for name, conditions := range testConditions {
		t.Run(name, func(t *testing.T) {
			sum := Product(conditions.nums)
			assert.Equal(t, conditions.expectedSum, sum)
		})
	}
}

func TestSum(t *testing.T) {
	testConditions := map[string]struct {
		nums []int
		expectedSum int
	} {
		"empty slice": {
			nums: []int{},
			expectedSum: 0,
		},
		"a few numbers": {
			nums: []int{ 1, 2, 3, 4, 5},
			expectedSum: 15,
		},
	}

	for name, conditions := range testConditions {
		t.Run(name, func(t *testing.T) {
			sum := Sum(conditions.nums)
			assert.Equal(t, conditions.expectedSum, sum)
		})
	}
}
