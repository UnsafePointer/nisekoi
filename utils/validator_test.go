package utils

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Validator", func() {
		g.It("should validate correct <github-org/repo> format", func() {
			g.Assert(ValidateSearchTerm("Ruenzuo/nisekoi")).Equal(true)
		})

		g.It("should validate correct <github-org> format", func() {
			g.Assert(ValidateSearchTerm("nisekoi")).Equal(true)
		})

		g.It("should validate incorrect <github-org/repo> format", func() {
			g.Assert(ValidateSearchTerm("Ruenzuo!#$/nisekoi#$@")).Equal(false)
		})

		g.It("should validate incorrect <github-org> format", func() {
			g.Assert(ValidateSearchTerm("nisekoi#$@")).Equal(false)
		})
	})
}
