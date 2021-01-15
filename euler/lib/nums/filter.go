package nums

// Filters down a given list of numbers to only include those numbers
// for which the given "FilterFunc" returns true
func Filter(nums []int, filterFunc func(num int) bool) []int {
	s := make([]int, 0)
	for _, num := range nums {
		if filterFunc(num) {
			s = append(s, num)
		}
	}

	return s
}
