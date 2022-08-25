package missing_letter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/TheMike316/codewars/challenges/missing_letter"
)

func dotest(s []rune, expected rune) {
	actual := FindMissingLetter(s)
	Expect(actual).To(Equal(expected),
		"With chars = %v (\"%s\")\nExpected %d ('%s') but got %d ('%s')", s, string(s), expected, string(expected), actual, string(actual))
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest([]rune{'a', 'b', 'c', 'd', 'f'}, 'e')
		dotest([]rune{'O', 'Q', 'R', 'S'}, 'P')
	})
})
