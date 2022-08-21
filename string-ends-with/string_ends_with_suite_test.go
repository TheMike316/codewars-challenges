package stringendswith_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringEndsWith(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringEndsWith Suite")
}
