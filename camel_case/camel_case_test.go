package camel_case_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/TheMike316/codewars/challenges/camel_case"

	. "fmt"
)

func dotest(str, exp string) {
	Println("input:", str)
	var ans = ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		dotest("", "")
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
		dotest("the-Stealth-Warrior", "theStealthWarrior")
	})
})
