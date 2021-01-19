package euler

import (
	"fmt"
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	result := SmallestMultiple()
	fmt.Printf("The smallest common multiple = %d\n", result)
}