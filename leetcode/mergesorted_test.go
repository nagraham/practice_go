package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoEqualLists(t *testing.T) {
	merged, err := merge([][]int{
		{ 1, 3 },
		{ 2, 4 },
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{ 1, 2, 3, 4 }, merged)
}

func TestMergeOneListEmpty(t *testing.T) {
	merged, err := merge([][]int{
		{},
		{ 1, 3, 5 },
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{ 1, 3, 5 }, merged)
}

func TestMergeOtherListEmpty(t *testing.T) {
	merged, err := merge([][]int{
		{ 1, 3, 5 },
		{},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{ 1, 3, 5 }, merged)
}

func TestMergeOneListFullyLessThanOther(t *testing.T) {
	merged, err := merge([][]int{
		{ 1, 3, 5 },
		{ 11, 13, 15 },
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{ 1, 3, 5, 11, 13, 15 }, merged)
}
