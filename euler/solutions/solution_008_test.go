package euler

import (
	"fmt"
	"testing"
)

func TestLargestProductInSeries(t *testing.T) {
	result := LargestProductInSeries()
	fmt.Printf("The largest product in the series = %d", result)
}