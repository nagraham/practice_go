package numbers

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

// Returns true if the number is even. This signature is of type Condition.
func IsEven(num int) bool {
	return num % 2 == 0
}