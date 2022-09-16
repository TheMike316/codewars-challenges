package twosum

func TwoSum(numbers []int, target int) [2]int {
	// at worst n(n-1)/2 performance; see handshake problem
	for i, num1 := range numbers {
		for j, num2 := range numbers[i+1:] {
			if num1+num2 == target {
				return [2]int{i, j + i + 1}
			}
		}
	}
	return [2]int{}
}
