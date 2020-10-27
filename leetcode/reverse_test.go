package leetcode


import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestReverseInt_Zero(t *testing.T) {
	result := reverseInt(0)
	assert.Equal(t, 0, result)
}

func TestReverseInt_One(t *testing.T) {
	result := reverseInt(1)
	assert.Equal(t, 1, result)
}

func TestReverseInt_NegativeOne(t *testing.T) {
	result := reverseInt(-1)
	assert.Equal(t, -1, result)
}

func TestReverse_HundredsPositive(t *testing.T) {
	result := reverseInt(123)
	assert.Equal(t, 321, result)
}

func TestReverse_HundredsNegative(t *testing.T) {
	result := reverseInt(-123)
	assert.Equal(t, -321, result)
}

func TestReverse_MaxInt32(t *testing.T) {
	result := reverseInt(math.MaxInt32)
	assert.Equal(t, 0, result)
}

func TestReverse_MinInt32(t *testing.T) {
	result := reverseInt(math.MinInt32)
	assert.Equal(t, 0, result)
}

var reversedInt int

func benchmarkReverseInt(i int, b *testing.B) {
	var result int
	for n := 0; n < b.N; n++ {
		result = reverseInt(i)
	}
	reversedInt = result
}

func BenchmarkReverseInt12345(b *testing.B) {
	benchmarkReverseInt(12345, b)
}

func BenchmarkReverseIntMaxInt(b *testing.B) {
	benchmarkReverseInt(math.MaxInt32, b)
}

// when 0
// when 1
// when -1
// when -123
// when Max int
// when min int