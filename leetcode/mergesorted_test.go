package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwo_EqualLists(t *testing.T) {
	merged, err := Merge([][]int{
		{1, 3},
		{2, 4},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, merged)
}

func TestMergeTwo_OneListEmpty(t *testing.T) {
	merged, err := Merge([][]int{
		{},
		{1, 3, 5},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{1, 3, 5}, merged)
}

func TestMergeTwo_OtherListEmpty(t *testing.T) {
	merged, err := Merge([][]int{
		{1, 3, 5},
		{},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{1, 3, 5}, merged)
}

func TestMergeTwo_ListFullyLessThanOther(t *testing.T) {
	merged, err := Merge([][]int{
		{1, 3, 5},
		{11, 13, 15},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{1, 3, 5, 11, 13, 15}, merged)
}

func TestMergeMultiple_MergesToSortedList(t *testing.T) {
	merged, err := Merge([][]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8},
		{0, 1, 1, 2, 3, 5, 8, 13},
		{-9, -6, -3, 0, 3, 6, 9},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{
		-9, -6, -3, 0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5, 6, 6, 7, 8, 8, 9, 9, 13,
	}, merged)
}

func TestMergeMultiple_HandlesEmptyLists(t *testing.T) {
	merged, err := Merge([][]int{
		{1, 2, 3},
		{},
		{4, 5, 6},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, []int{
		1, 2, 3, 4, 5, 6,
	}, merged)
}
