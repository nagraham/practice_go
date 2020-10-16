package leetcode

import (
	"errors"
	"log"
	"math"
)

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// As a followup, solve in O(log (n+m)) <- this is very difficult, I spent days on this last year

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArraysTwo(nums1, nums2)
}

/*
 * Approach One: O(n/2)
 * 1) Sum lengths of lists; divide that in two -> target; if odd, look for 1; if even get 2
 * 2) Keep a count (if odd, set to 0; if even, set to 1)
 * 3) Walk lists with pointers, one at a time, increment count, incrementing the pointer at the lesser value
 * 4) When count == target, get 1 or 2 numbers, and return median
 *    Getting 2 numbers, you need to increment the least pointer, and check
 *
 * Edge cases:
 *  - One list much longer than the other
*   - Even lists much longer than other (both values from same list)
 */
func findMedianSortedArraysOne(nums1 []int, nums2 []int) float64 {
	combinedLen := len(nums1) + len(nums2)

	midPoint := combinedLen / 2
	if isEven(combinedLen) {
		midPoint -= 1
	}

	indexPtr1, indexPtr2 := walkLists(midPoint, nums1, nums2)

	val1 := extractValue(nums1, indexPtr1, math.MaxInt64)
	val2 := extractValue(nums2, indexPtr2, math.MaxInt64)

	// if even, we need to determine which two values are lesser (they may be from the same list)
	if isEven(combinedLen) {
		var a int
		if val1 < val2 {
			a = val1
			val1 = extractValue(nums1, indexPtr1 + 1, math.MaxInt64)
		} else {
			a = val2
			val2 = extractValue(nums2, indexPtr2 + 1, math.MaxInt64)
		}

		if val1 < val2 {
			return (float64(a) + float64(val1)) / 2
		} else {
			return (float64(a) + float64(val2)) / 2
		}

	} else {
		if val1 < val2 {
			return float64(val1)
		} else {
			return float64(val2)
		}
	}
}

// Walk the sorted lists to a target index; should be 0 <= target < len(nums1 + nums2)
// Return two pointers to indices for both lists when the target is reached;
// returns the pointers in the same order as the provided lists.
func walkLists(target int, nums1 []int, nums2 []int) (int, int) {
	var count int

	indexPtr1 := 0
	indexPtr2 := 0

	// Walk the sorted lists with two pointers until we arrive at a virtual mid point
	for count < target {
		val1 := extractValue(nums1, indexPtr1, math.MaxInt64)
		val2 := extractValue(nums2, indexPtr2, math.MaxInt64)

		if val1 < val2 {
			indexPtr1++
		} else {
			indexPtr2++
		}

		count++
	}

	return indexPtr1, indexPtr2
}

func isEven(n int) bool {
	return n % 2 == 0
}

func extractValue(nums []int, index int, defaultVal int) int {
	if index < len(nums) {
		return nums[index]
	} else {
		return defaultVal
	}
}


/*
 * Approach Two: Compose sub-problems
 *
 * 1) Implement function to merge 2 sorted lists
 * 2) Implement function to get median value of sorted list
 *
 * This approach is O(N), but results in orthogonal code that is re-usable in other solutions, and
 * is going to be easier to maintain. I wish I had started with this approach. The problem is I
 * started thinking of Approach #1 and got "hooked" into puzzling it out.
 */

// Find median number in a sorted list
func median(nums []int) (float64, error) {
	isEven := len(nums) % 2 == 0
	midpoint := len(nums) / 2

	if len(nums) == 0 {
		return 0.0, errors.New("given list cannot be empty")
	} else if isEven {
		return (float64(nums[midpoint - 1]) + float64(nums[midpoint])) / 2, nil
	} else {
		return float64(nums[midpoint]), nil
	}
}

func findMedianSortedArraysTwo(nums1 []int, nums2 []int) float64 {
	merged, err := merge([][]int{ nums1, nums2 })
	if err != nil {
		log.Panic(err)
	}

	m, err := median(merged)
	if err != nil {
		log.Panic(err)
	}

	return m
}