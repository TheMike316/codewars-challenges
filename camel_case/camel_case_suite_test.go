package camel_case_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCamelCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CamelCase Suite")
}
