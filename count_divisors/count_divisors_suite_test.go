package count_divisors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCountDivisors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountDivisors Suite")
}
