package euler

import (
	"euler/lib/numbers"
)

func LargestPrimeFactor() int {
	factors := numbers.Factors(600851475143)
	primeFactors := numbers.Filter(factors, numbers.IsPrime)
	return primeFactors[len(primeFactors)-1]
}
