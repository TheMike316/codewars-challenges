package highest_scoring_word_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHighestScoringWord(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HighestScoringWord Suite")
}
