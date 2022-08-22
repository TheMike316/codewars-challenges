package population_growth_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pg "github.com/TheMike316/codewars/challenges/population_growth"
)

var _ = Describe("PopulationGrowth", func() {
	It("fixed tests", func() {
		Expect(pg.NbYear(1500, 5, 100, 5000)).To(Equal(15))
		Expect(pg.NbYear(1500000, 2.5, 10000, 2000000)).To(Equal(10))
		Expect(pg.NbYear(1500000, 0.25, 1000, 2000000)).To(Equal(94))
		Expect(pg.NbYear(1500000, 0.25, -1000, 2000000)).To(Equal(151))
	  })
})
