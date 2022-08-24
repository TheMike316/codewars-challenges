package roman_numerals_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/TheMike316/codewars/challenges/roman_numerals"
)

var _ = Describe("RomanNumerals", func() {
	It("should give decimal number from roman", func() {
		Expect(Decode("XXI")).To(Equal(21))

	})
	It("should give decimal number from roman", func() {
		Expect(Decode("I")).To(Equal(1))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("IV")).To(Equal(4))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MMVIII")).To(Equal(2008))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MDCLXVI")).To(Equal(1666))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MCDX")).To(Equal(1410))
	})
})
