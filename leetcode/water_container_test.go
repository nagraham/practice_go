package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateWaterArea(t *testing.T) {
	testContexts := map[string]struct {
		heights []int
		i int
		j int
		expectedArea int
	} {
		"with two elements, one shorter": {
			[]int{ 4, 3 }, 0, 1, 3,
		},
		"with multiple elements": {
			[]int{ 4, 1, 3 }, 0, 2, 6,
		},
		"with zero": {
			[]int{ 0, 1, 3 }, 0, 2, 0,
		},
		"with many elements": {
			[]int{ 1, 8, 6, 2, 5, 4, 8, 3, 7 }, 1, 8, 49,
		},
	}

	for context, data := range testContexts {
		t.Run(context, func (t *testing.T){
			area := calculateWaterArea(data.heights, data.i, data.j)
			assert.Equal(t, area, data.expectedArea)
		})
	}
}

func TestMaxWaterContainer(t *testing.T) {
	testContexts := map[string]struct {
		heights []int
		expectedMaxArea int
	} {
		"with two ones": {
			[]int{ 1, 1 }, 1,
		},
		"with three ones": {
			[]int{ 1, 1, 1 }, 2,
		},
		"with four ones": {
			[]int{ 1, 1, 1, 1}, 3,
		},
		"with ones and a single two": {
			[]int{ 1, 1, 2, 1}, 3,
		},
		"with ones and twos": {
			[]int{ 1, 2, 2, 1}, 3,
		},
		"with many elements": {
			[]int{ 1, 8, 6, 2, 5, 4, 8, 3, 7 }, 49,
		},
		"with many elements and a set of very tall values": {
			[]int{ 1, 8, 6, 2, 5, 4, 50, 50, 7 }, 50,
		},
		"with many elements and a set of very tall values closer to a good answer": {
			[]int{ 1, 8, 50, 50, 5, 4, 8, 3, 7 }, 50,
		},
	}

	for context, data := range testContexts {
		t.Run(context, func (t *testing.T){
			area := MaxWaterContainer(data.heights)
			assert.Equal(t, data.expectedMaxArea, area)
		})
	}
}
