package weird_case_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	wc "github.com/TheMike316/codewars/challenges/weird_case"
)

var _ = Describe("WeirdCase", func() {
	It("Should return the correct values", func() {
		Expect(wc.ToWeirdCase("abc def")).To(Equal("AbC DeF"))
		Expect(wc.ToWeirdCase("ABC")).To(Equal("AbC"))
		Expect(wc.ToWeirdCase("This is a test Looks like you passed")).To(Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
	})
})
