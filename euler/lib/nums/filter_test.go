package nums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	testConditions := map[string]struct {
		nums []int
		filterFunc func(num int) bool
		expectedValues []int
	} {
		"filter to even numbers": {
			nums: []int{ 1, 2, 3, 4, 5},
			filterFunc: func(n int) bool {
				return n % 2 == 0
			},
			expectedValues: []int{ 2, 4 },
		},
		"filter to odd numbers": {
			nums: []int{ 1, 2, 3, 4, 5},
			filterFunc: func(n int) bool {
				return n % 2 == 1
			},
			expectedValues: []int{ 1, 3, 5 },
		},
		"filter to multiples of five": {
			nums: []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			filterFunc: func(n int) bool {
				return n % 5 == 0
			},
			expectedValues: []int{ 5, 10 },
		},
	}

	for name, conditions := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := Filter(conditions.nums, conditions.filterFunc)
			assert.Equal(t, conditions.expectedValues, result)
		})
	}

}