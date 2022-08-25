package missing_letter

func FindMissingLetter(chars []rune) rune {
	// we can assume len(chars) > 1 and always in a correct sequence with only one letter missing
	for i := 1; i < len(chars); i++ {
		r := chars[i]
		diff := r - chars[i-1]
		if diff > 1 {
			return rune(int(r) - 1)
		}
	}
	panic("illegal argument")
}
