package utils

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := Goblin(t)

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Validator", func() {
		g.It("should validate correct <github-org/repo> format", func() {
			_, _, err := ValidateSearchTerm("Ruenzuo/nisekoi")
			Expect(err).To(BeNil())
		})

		g.It("should validate correct <github-org> format", func() {
			_, _, err := ValidateSearchTerm("nisekoi")
			Expect(err).To(BeNil())
		})

		g.It("should validate incorrect <github-org/repo> format", func() {
			_, _, err := ValidateSearchTerm("Ruenzuo!#$/nisekoi#$@")
			Expect(err).ToNot(BeNil())
		})

		g.It("should validate incorrect <github-org> format", func() {
			_, _, err := ValidateSearchTerm("nisekoi#$@")
			Expect(err).ToNot(BeNil())
		})
	})
}
