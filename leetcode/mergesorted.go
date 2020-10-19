package leetcode

import (
	"datastructure"
)

// Merge two or more sorted lists of integers
func Merge(numsList [][]int) ([]int, error) {
	if len(numsList) > 2 {
		return mergeMultiple(numsList), nil
	} else if len(numsList) == 2 {
		return mergeTwo(numsList[0], numsList[1]), nil
	} else {
		return numsList[0], nil
	}
}

// NOTE: In production code, I would probably only keep the `mergeMultiple` function ... but
// given that this is practice code, I'll keep this one around
func mergeTwo(nums1 []int, nums2 []int) []int {
	merged := make([]int, len(nums1)+len(nums2))

	index1, index2 := 0, 0

	for count := 0; count < len(nums1)+len(nums2); count++ {
		isOneDone := index1 == len(nums1)
		isTwoDone := index2 == len(nums2)

		if isOneDone {
			merged[count] = nums2[index2]
			index2++

		} else if isTwoDone {
			merged[count] = nums1[index1]
			index1++

		} else {
			if nums1[index1] < nums2[index2] {
				merged[count] = nums1[index1]
				index1++
			} else {
				merged[count] = nums2[index2]
				index2++
			}
		}
	}

	return merged
}

type listTraversalState struct {
	list      []int
	currIndex int
}

func mergeMultiple(listOfLists [][]int) []int {

	// Set up priority queue.
	pq := datastructure.NewPriorityQueue()
	size := 0
	for _, list := range listOfLists {
		if len(list) > 0 {
			pq.Push(&listTraversalState{list: list, currIndex: 0}, list[0])
			size += len(list)
		}
	}

	// Walk each list, adding values to a single merged list, one at a time, using the
	// priority queue to always choose the lowest number from any list.
	merged := make([]int, size)
	for i := 0; pq.Len() > 0; i++ {
		listPtr := pq.Pop().(*listTraversalState)
		merged[i] = listPtr.list[listPtr.currIndex]
		listPtr.currIndex = listPtr.currIndex + 1
		if listPtr.currIndex < len(listPtr.list) {
			pq.Push(listPtr, listPtr.list[listPtr.currIndex])
		}
	}
	return merged
}
