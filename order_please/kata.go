package order_please

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type wordSplit []string

var rp = regexp.MustCompile("[0-9]")

func Order(sentence string) string {
	if sentence == "" {
		return sentence
	}

	words := strings.Split(sentence, " ")
	sort.Sort(wordSplit(words))

	return strings.Join(words, " ")
}

func (s wordSplit) Len() int {
	return len(s)
}

func (s wordSplit) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s wordSplit) Less(i, j int) bool {
	return getWordPos(s[i]) < getWordPos(s[j])
}

func getWordPos(str string) int {
	if i, err := strconv.Atoi(rp.FindAllString(str, 1)[0]); err == nil {
		return i
	} else {
		return -1
	}
}
