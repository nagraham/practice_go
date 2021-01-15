package fibonacci

// Returns a Fibonnaci sequence made up of values no greater than the given max
func UpTo(max int) []int {
	if max == 0 {
		return []int{}
	}
	if max == 1 {
		return []int{0}
	}

	fibs := []int{0, 1}

	for next(fibs) < max {
		fibs = append(fibs, next(fibs))
	}

	return fibs
}

func next(fibs []int) int {
	return fibs[len(fibs) - 1] + fibs[len(fibs) - 2]
}