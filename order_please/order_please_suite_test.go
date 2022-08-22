package order_please_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrderPlease(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderPlease Suite")
}
