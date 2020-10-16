package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrayShortAndOdd(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 3}, []int{2})
	assert.Equal(t, 2.0, median)
}

func TestFindMedianSortedArrayShortAndEven(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	assert.Equal(t, 2.5, median)
}

func TestFindMedianSortedArrayZeros(t *testing.T) {
	median := findMedianSortedArrays([]int{0, 0}, []int{0, 0})
	assert.Equal(t, 0.0, median)

	median = findMedianSortedArrays([]int{0}, []int{0, 0})
	assert.Equal(t, 0.0, median)
}

func TestFindMedianSortedArrayUnevenOddLength(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9})
	assert.Equal(t, 5.0, median)

	median = findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6, 7}, []int{8, 9})
	assert.Equal(t, 5.0, median)
}

func TestFindMedianSortedArrayUnevenEvenLength(t *testing.T) {
	median := findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 5.5, median)

	median = findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{9, 10})
	assert.Equal(t, 5.5, median)
}

func TestMedianEmptyList(t *testing.T) {
	_, err := median([]int{})
	assert.Error(t, err)
}

func TestMedianOneElement(t *testing.T) {
	m, err := median([]int{ 42 })
	if err != nil {
		t.Fatal(m)
	}
	assert.Equal(t, 42.0, m)
}

func TestMedianTwoElements(t *testing.T) {
	m, err := median([]int{ 2, 2 })
	if err != nil {
		t.Fatal(m)
	}
	assert.Equal(t, 2.0, m)
}

func TestMedianMultipleElementsEven(t *testing.T) {
	m, err := median([]int{ 1, 2, 3, 4, 5, 6 })
	if err != nil {
		t.Fatal(m)
	}
	assert.Equal(t, 3.5, m)
}

func TestMedianMultipleElementsOdd(t *testing.T) {
	m, err := median([]int{ 1, 2, 3, 4, 5 })
	if err != nil {
		t.Fatal(m)
	}
	assert.Equal(t, 3.0, m)
}