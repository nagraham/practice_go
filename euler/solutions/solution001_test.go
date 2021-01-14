package euler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumMultiplesOfThreeUnderTen(t *testing.T) {
	sum := SumMultiples(10, 3)
	assert.Equal(t, 18, sum)
}

func TestSumMultiplesOfTwoUnderTen(t *testing.T) {
	sum := SumMultiples(10, 2)
	assert.Equal(t, 20, sum)
}

// 2, 3, 4, 6, 8, 9 (do not double count)
func TestSumMultiplesOfTwoAndThreeUnderTen(t *testing.T) {
	sum := SumMultiples(10, 2, 3)
	assert.Equal(t, 32, sum)
}


func TestSumMultiplesOfThreeAndFiveUnderTen(t *testing.T) {
	sum := SumMultiples(10, 3, 5)
	assert.Equal(t, 23, sum)
}

func TestSumMultiplesOfThreeAndFiveUnder1000(t *testing.T) {
	sum := SumMultiples(1000, 3, 5)
	// fmt.Printf("Answer = %d", sum)
	assert.Equal(t, 233168, sum)
}