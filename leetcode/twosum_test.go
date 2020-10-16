package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertIndices(t *testing.T, result []int, i int, j int) {
	assert.Contains(t, result, i)
	assert.Contains(t, result, j)
}

func TestTwoSumPositiveNums(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	assertIndices(t, result, 0, 1)
}

func TestTwoSumCandidateNeedsDuplicateNotPresent(t *testing.T) {
	// Candidate(3): 6 - 3 = 3, but only 1 is present in list
	result := twoSum([]int{3, 2, 4}, 6)
	assertIndices(t, result, 1, 2)
}

func TestTwoSumCandidateNeedsDuplicatePresent(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	assertIndices(t, result, 0, 1)
}

func TestTwoSumNegativeNumbers(t *testing.T) {
	result := twoSum([]int{-7, 3, -2, 0, -3}, -9)
	assertIndices(t, result, 0, 2)
}

func TestTwoSumNegativeNumbersOnePositiveOneNegative(t *testing.T) {
	result := twoSum([]int{7, 3, -1, 0, -3, -16}, -9)
	assertIndices(t, result, 0, 5)
}
