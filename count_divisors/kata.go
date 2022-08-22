package count_divisors

import "math"

func Divisors(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	// 1 and itself
	count := 2
	for i := 2; i <= n/2; i++ {
		if math.Mod(float64(n), float64(i)) == 0 {
			count++
		}
	}
	return count
}
