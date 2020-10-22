package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// with 1, with 4, with empty string, with string length 1

func TestZigZag_withStringAndThreeRows(t *testing.T) {
	result := zigzag("PAYPALISHIRING", 3)
	assert.Equal(t, "PAHNAPLSIIGYIR", result)
}

func TestZigZag_withStringAndOneRow(t *testing.T) {
	result := zigzag("PAYPALISHIRING", 1)
	assert.Equal(t, "PAYPALISHIRING", result)
}

func TestZigZag_withStringAndFourRows(t *testing.T) {
	result := zigzag("PAYPALISHIRING", 4)
	assert.Equal(t, "PINALSIGYAHRPI", result)
}

func TestZigZag_withBlankStringAndFourRows(t *testing.T) {
	result := zigzag("", 4)
	assert.Equal(t, "", result)
}

func TestZigZagOscillator(t *testing.T) {
	zz := NewZigZagOscillator(0, 5)

	values := make([]int, 15)
	for i := 0; i < 15; i++ {
		values[i] = zz.Get()
		zz.Move()
	}

	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 3, 2, 1, 0, 1, 2, 3, 4, 3, 2}, values)
}

func TestZigZagOscillator_OscillateAroundOneValue(t *testing.T) {
	zz := NewZigZagOscillator(0, 1)

	values := make([]int, 10)
	for i := 0; i < 10; i++ {
		values[i] = zz.Get()
		zz.Move()
	}

	assert.ElementsMatch(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, values)
}

