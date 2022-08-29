package camel_case

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile("[-_]")
	words := re.Split(s, -1)
	if len(words) == 1 {
		return s
	}
	for i := 1; i < len(words); i++ {
		word := words[i]
		c := strings.ToUpper(string(word[0]))
		words[i] = c + word[1:]
	}
	return strings.Join(words, "")
}
