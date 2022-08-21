package is_prime_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsPrime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsPrime Suite")
}
