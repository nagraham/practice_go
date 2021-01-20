package numbers

// Returns true if the number is even. This signature is of type Condition.
func IsEven(num int) bool {
	return num % 2 == 0
}

// Returns the largest of the two integers
func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Returns the smallest of the two integers
func MinInt(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// Returns a new list with the given range of integers, between "start" and "end" (exclusive).
// If "start" is greater or equal to "end" returns an empty list.
func Range(start int, end int) []int {
	if start >= end {
		return []int{}
	}

	length := end-start
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = start
		start++
	}

	return nums
}

