package numbers

// Returns a new list with the contents reversed
func Reverse(nums []int) []int {
	reversed := make([]int, len(nums))

	for start, end := 0, len(nums)-1; start <= end; start, end = start+1, end-1 {
		reversed[start] = nums[end]
		reversed[end] = nums[start]
	}

	return reversed
}