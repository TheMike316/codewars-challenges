package order_please_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	order "github.com/TheMike316/codewars/challenges/order_please"
)

var _ = Describe("OrderPlease", func() {
	It("Sample tests", func() {
		Expect(order.Order("is2 Thi1s T4est 3a")).To(Equal("Thi1s is2 3a T4est"))
		Expect(order.Order("4of Fo1r pe6ople g3ood th5e the2")).To(Equal("Fo1r the2 g3ood 4of th5e pe6ople"))
		Expect(order.Order("")).To(Equal(""))
	})
})
