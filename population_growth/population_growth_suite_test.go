package population_growth_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPopulationGrowth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PopulationGrowth Suite")
}
