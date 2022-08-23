package pile_of_cubes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPileOfCubes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PileOfCubes Suite")
}
