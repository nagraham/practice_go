package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpTo(t *testing.T) {

	conditions := map[string]struct {
		max int
		expectedValues []int
	} {
		"When zero": {
			max: 0,
			expectedValues: []int{ },
		},
		"When one": {
			max: 1,
			expectedValues: []int{ 0 },
		},
		"When two": {
			max: 2,
			expectedValues: []int{ 0, 1, 1 },
		},
		"When up to ten": {
			max: 10,
			expectedValues: []int{ 0, 1, 1, 2, 3, 5, 8 },
		},
		"When up to thirty-four": {
			max: 34,
			expectedValues: []int{ 0, 1, 1, 2, 3, 5, 8, 13, 21 },
		},
		"When up to thirty-five": {
			max: 35,
			expectedValues: []int{ 0, 1, 1, 2, 3, 5, 8, 13, 21, 34 },
		},
	}

	for name, cond := range conditions {
		t.Run(name, func(t *testing.T) {
			fibs := UpTo(cond.max)
			assert.Equal(t, cond.expectedValues, fibs)
		});
	}
}