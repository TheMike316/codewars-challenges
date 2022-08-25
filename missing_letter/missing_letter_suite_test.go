package missing_letter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMissingLetter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MissingLetter Suite")
}
