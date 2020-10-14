package leetcode

// Return the indices of two different numbers from the given array
// such that they add to the given target.
func twoSum(nums []int, target int) []int {
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