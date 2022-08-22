package count_divisors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	div "github.com/TheMike316/codewars/challenges/count_divisors"
)

var _ = Describe("CountDivisors", func() {
	It("Divisors(1)", func() {Expect(div.Divisors(1)).To(Equal(1))})
	It("Divisors(10)", func() {Expect(div.Divisors(10)).To(Equal(4))})
	It("Divisors(11)", func() {Expect(div.Divisors(11)).To(Equal(2))})
	It("Divisors(54)", func() {Expect(div.Divisors(54)).To(Equal(8))})
	It("Divisors(64)", func() {Expect(div.Divisors(64)).To(Equal(7))})
})
