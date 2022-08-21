package is_prime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/TheMike316/codewars/challenges/is_prime"
)

var _ = Describe("IsPrime", func() {
	It("Should return true on primes, false on other numbers", func() {
		Expect(is_prime.IsPrime(1)).To(Equal(false))
		Expect(is_prime.IsPrime(2)).To(Equal(true))
		Expect(is_prime.IsPrime(-1)).To(Equal(false))
		Expect(is_prime.IsPrime(999)).To(Equal(false))
		Expect(is_prime.IsPrime(13)).To(Equal(true))
		Expect(is_prime.IsPrime(9971207)).To(Equal(true))
		Expect(is_prime.IsPrime(62114581)).To(Equal(true))
	})
})
