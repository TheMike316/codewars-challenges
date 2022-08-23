package pile_of_cubes

func FindNb(m int) int {
	volume := 0
	for n := 1; volume < m; n++ {
		volume += n * n * n
		if volume == m {
			return n
		}
	}
	return -1
}
