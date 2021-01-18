package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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