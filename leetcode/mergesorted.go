package leetcode

import "errors"

// Merge two sorted lists of integers
func merge(numsList [][]int) ([]int, error) {
	if len(numsList) > 2 {
		return nil, errors.New("more than 2 lists not yet supported")
	} else if len(numsList) == 2 {
		return mergeTwo(numsList[0], numsList[1]), nil
	} else {
		return numsList[0], nil
	}
}

func mergeTwo(nums1 []int, nums2 []int) []int {
	merged := make([]int, len(nums1) + len(nums2))

	index1, index2 := 0, 0

	for count := 0; count < len(nums1) + len(nums2); count++ {
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