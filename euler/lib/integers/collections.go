package numbers

// A function that takes an int
type IntFunction = func(n int) int

// A function that evaluates a number and returns true or false based on
// its condition expressions.
type Condition = func(num int) bool

// Filters down a given list of numbers to only include those numbers
// for which the given Condition returns true
func Filter(nums []int, filter Condition) []int {
	s := make([]int, 0)
	for _, num := range nums {
		if filter(num) {
			s = append(s, num)
		}
	}

	return s
}

// A Mapping function that loops through each number, applies the given mapper function,
// and returns a list of those results.
func Map(nums []int, mapper IntFunction) []int {
	mappedNums := make([]int, len(nums))
	for i, n := range nums {
		mappedNums[i] = mapper(n)
	}
	return mappedNums
}

// Returns a new list with the contents reversed
func Reverse(nums []int) []int {
	reversed := make([]int, len(nums))

	for start, end := 0, len(nums)-1; start <= end; start, end = start+1, end-1 {
		reversed[start] = nums[end]
		reversed[end] = nums[start]
	}

	return reversed
}