package pile_of_cubes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/TheMike316/codewars/challenges/pile_of_cubes"
)

func dotest(n int, exp int) {
	var ans = FindNb(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(4183059834009, 2022)
		dotest(24723578342962, -1)
	})
})
