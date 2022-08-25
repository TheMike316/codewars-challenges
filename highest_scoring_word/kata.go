package highest_scoring_word

import "strings"

func High(s string) string {
	highScore := 0
	highest := ""
	for _, word := range strings.Split(s, " ") {
		score := 0
		for _, r := range word {
			score += int(r) + 1 - int('a')
		}
		if score > highScore {
			highest = word
			highScore = score
		}
	}
	return highest
}
