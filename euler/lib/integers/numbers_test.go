package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	testConditions := map[string]struct {
		start int
		end int
		expectedResult []int
	} {
		"end is less than start": {
			start: 3,
			end: 2,
			expectedResult: emptyList,
		},
		"end is equal to start": {
			start: 3,
			end: 3,
			expectedResult: emptyList,
		},
		"end is start+1": {
			start: 1,
			end: 2,
			expectedResult: []int{ 1 },
		},
		"end is start+5": {
			start: 4,
			end: 9,
			expectedResult: []int{ 4, 5, 6, 7, 8 },
		},
	}

	for name, cond := range testConditions {
		t.Run(name, func(t *testing.T) {
			result := Range(cond.start, cond.end)
			assert.Equal(t, cond.expectedResult, result)
		})
	}
}
