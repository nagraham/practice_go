package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestMyAtoi_PositiveNumber(t *testing.T) {
	result := myAtoi("1234567890")
	assert.Equal(t, 1234567890, result)
}

func TestMyAtoi_Zero(t *testing.T) {
	result := myAtoi("0")
	assert.Equal(t, 0, result)
}

func TestMyAtoi_ManyZeroes(t *testing.T) {
	result := myAtoi("00000")
	assert.Equal(t, 0, result)
}

func TestMyAtoi_LeadingZeroesAreDiscarded(t *testing.T) {
	result := myAtoi("000001")
	assert.Equal(t, 1, result)
}

func TestMyAtoi_ALotOfLeadingZeros(t *testing.T) {
	result := myAtoi("  0000000000012345678")
	assert.Equal(t, 12345678, result)
}

func TestMyAtoi_leadingWhitespaceIsTrimmed(t *testing.T) {
	result := myAtoi("   123")
	assert.Equal(t, 123, result)
}

func TestMyAtoi_onlyWhitespaceReturnsZero(t *testing.T) {
	result := myAtoi(" ")
	assert.Equal(t, 0, result)
}

func TestMyAtoi_positiveSignReturnsPositiveNumber(t *testing.T) {
	result := myAtoi("+123")
	assert.Equal(t, 123, result)
}

func TestMyAtoi_negativeSignReturnsNegativeNumber(t *testing.T) {
	result := myAtoi("-123")
	assert.Equal(t, -123, result)
}

func TestMyAtoi_TrailingWhiteSpaceIsTrimmed(t *testing.T) {
	result := myAtoi("123   ")
	assert.Equal(t, 123, result)
}

func TestMyAtoi_NonNumberCharactersCauseReturn(t *testing.T) {
	result := myAtoi("1a23")
	assert.Equal(t, 1, result)
}

func TestMyAtoi_ManyNonNumberCharactersCauseReturn(t *testing.T) {
	result := myAtoi("   123   hello world!")
	assert.Equal(t, 123, result)
}

func TestMyAtoi_NumberGreaterThanMaxIntReturnsMaxInt(t *testing.T) {
	result := myAtoi(strconv.Itoa(math.MaxInt32 + 1))
	assert.Equal(t, math.MaxInt32, result)
}

func TestMyAtoi_MaxInt(t *testing.T) {
	result := myAtoi(strconv.Itoa(math.MaxInt32))
	assert.Equal(t, math.MaxInt32, result)
}

func TestMyAtoi_NumberLessThanMinIntReturnsMinInt(t *testing.T) {
	result := myAtoi(strconv.Itoa(math.MinInt32 - 1))
	assert.Equal(t, math.MinInt32, result)
}

func TestMyAtoi_MinInt(t *testing.T) {
	result := myAtoi(strconv.Itoa(math.MinInt32))
	assert.Equal(t, math.MinInt32, result)
}

func TestMyAtoi_VeryLargeNumber(t *testing.T) {
	result := myAtoi("9223372036854775808")
	assert.Equal(t, math.MaxInt32, result)
}

func TestMyAtoi_VeryLargeNegativeNumber(t *testing.T) {
	result := myAtoi("-19223372036854775808")
	assert.Equal(t, math.MinInt32, result)
}
