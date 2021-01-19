package numbers

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
