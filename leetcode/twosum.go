package leetcode

func twoSum(nums []int, target int) []int {
	return twoSumOne(nums, target)
}


// Return the indices of two different numbers from the given array such that they add to the
// given target.
//
// Given that (a + b) = c, given 'a' and 'c' we can find 'b' with (c - a). We can use math to get
// a possible 'b' and then search for it. You could iterate through the array, but I prefer using
// a map. However, we need to capture the possibility of duplicates. Especially for the case that
// c / a == b. So our map can store a list of indicies.
func twoSumOne(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var numMap map[int][]int
	numMap = mapByIndex(nums)

	for i, num := range nums {
		indices, hasVal := numMap[target - num]
		if hasVal {
			j, isOtherIndex := someOtherNumber(i, indices)
			if isOtherIndex {
				return []int{ i, j }
			}
		}
	}

	return []int{}
}

// Given slice of nums: return map with key=number, value=slice of indices
func mapByIndex(nums[] int) map[int][]int {
	numMap := make(map[int][]int)
	for i, num := range nums {
		s, hasValue := numMap[num]
		if hasValue {
			numMap[num] = append(s, i)
		} else {
			numMap[num] = []int{ i }
		}
	}
	return numMap
}

// Searches slice for some other number and returns that number and an "ok" boolean indicating
// whether the number is different. If not different, returns the given num and false.
func someOtherNumber(num int, s []int) (int, bool) {
	for _, other := range s {
		if num != other {
			return other, true
		}
	}
	return num, false
}

// ATTEMPT TWO:
// I saw an algorithm here in the Leetcode forum, and I wanted to re-write it to burn the
// simplicity of it into my mind. It's nice and concise (but I believe my solution is
// better suited to extending to ThreeSum
func twoSumTwo(nums []int, target int) []int {
	numMap := make(map[int]int)

	// b/c we loop once, only duplicate values will be present in the map.
	for index, num := range nums {
		otherIndex, hasNum := numMap[target - num]
		if hasNum {
			return []int{ index, otherIndex }
		}
		numMap[num] = index
	}

	return []int{}
}