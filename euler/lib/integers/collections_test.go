package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var emptyList = []int{}


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

func TestMap(t *testing.T) {
	nums := []int{ 1, 2, 3, 4, 5 }
	result := Map(nums, func(n int) int {
		return n + n
	})
	expected := []int{ 2, 4, 6, 8, 10 }
	assert.Equal(t, expected, result)
}

func TestReverse(t *testing.T) {
	testConditions := map[string]struct {
		nums []int
		expectedResult []int
	} {
		"empty list": {
			nums: []int{},
			expectedResult: []int{},
		},
		"when size=1": {
			nums: []int{1},
			expectedResult: []int{1},
		},
		"when size=2": {
			nums: []int{1,2},
			expectedResult: []int{2,1},
		},
		"when size > 2 and odd": {
			nums: []int{1,2,3},
			expectedResult: []int{3,2,1},
		},
		"when size > 2 and even": {
			nums: []int{1,2,3,4},
			expectedResult: []int{4,3,2,1},
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := Reverse(cond.nums)
			assert.Equal(t, cond.expectedResult, result)
		})
	}
}