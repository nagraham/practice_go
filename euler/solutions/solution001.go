package euler

// Returns the sum of all multiples of each provided num up to the max
func SumMultiples(max int, nums... int) int {
	sum := 0

	for i := 1; i < max; i++ {
		for _, num := range nums {
			if i % num == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
