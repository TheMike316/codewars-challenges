package roman_numerals

func Decode(roman string) int {
	numerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(roman); i++ {
		r := roman[i]
		// we can assume the ruman numeral is in a correct format
		if r == 'I' || r == 'X' || r == 'C' {
			if i+1 < len(roman) {
				n := numerals[rune(r)]
				n2 := numerals[rune(roman[i+1])]
				if n2 > n {
					result += n2 - n
					i++
					continue
				}
			}
		}
		result += numerals[rune(r)]
	}
	return result
}
