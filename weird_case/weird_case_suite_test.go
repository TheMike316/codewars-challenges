package weird_case_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWeirdCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WeirdCase Suite")
}
