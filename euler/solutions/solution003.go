package euler

import (
	"euler/lib/integers"
)

func LargestPrimeFactor() int {
	factors := integers.Factors(600851475143)
	primeFactors := integers.Filter(factors, integers.IsPrime)
	return primeFactors[len(primeFactors)-1]
}
