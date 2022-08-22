package weird_case

import (
	"math"
	"strings"
)

func ToWeirdCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		weirdWord := ""
		for j, s := range word {
			if math.Mod(float64(j), 2) == 0 {
				weirdWord += strings.ToUpper(string(s))
			} else {
				weirdWord += strings.ToLower(string(s))
			}
		}
		words[i] = weirdWord
	}
	return strings.Join(words, " ")
}
