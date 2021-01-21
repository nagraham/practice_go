package integers

// Get product from a slice of integers.
func Product(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}

// Sum up a slice of integers.
func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
