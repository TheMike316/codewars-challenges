package is_prime

import (
	"math"
)

func IsPrime(n int) bool {
	// Numbers go up to 2^31
	if n <= 1 {
		return false
	}

	// no fancypants algo here, test ints from 2 to sqrt of n
	limit := math.Sqrt(float64(n))
	for i := 2; float64(i) <= limit; i++ {
		if math.Mod(float64(n), float64(i)) == 0 {
			return false
		}
	}

	return true
}
