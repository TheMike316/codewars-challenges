package population_growth

import "math"

func NbYear(p0 int, percent float64, aug int, p int) int {
	growthFactor := percent / 100
	pn := p0
	var n int
	for n = 0; pn < p; n++ {
		pn = pn + int(math.Floor(float64(pn)*growthFactor)) + aug
	}
	return n
}
