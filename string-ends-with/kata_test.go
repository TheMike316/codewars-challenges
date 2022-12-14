package stringendswith

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func solution(str, ending string) bool {
	// boring answer
	return strings.HasSuffix(str, ending)
	// slightly less boring answer
	// return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}

var _ = Describe("Example test:", func() {
	It("Should work for fixed tests", func() {
		Expect(solution("", "")).To(Equal(true))
		Expect(solution(" ", "")).To(Equal(true))
		Expect(solution("abc", "c")).To(Equal(true))
		Expect(solution("banana", "ana")).To(Equal(true))
		Expect(solution("a", "z")).To(Equal(false))
		Expect(solution("", "t")).To(Equal(false))
		Expect(solution("$a = $b + 1", "+1")).To(Equal(false))
		Expect(solution("    ", "   ")).To(Equal(true))
		Expect(solution("1oo", "100")).To(Equal(false))
	})
})
