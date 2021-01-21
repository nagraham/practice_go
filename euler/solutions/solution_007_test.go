package euler

import (
	"fmt"
	"testing"
)

func TestTenThousandthAndOnePrime(t *testing.T) {
	result := TenThousandthAndOnePrime()
	fmt.Printf("The 10001th prime is = %d\n", result)
}