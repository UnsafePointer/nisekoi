package utils

import (
	"regexp"
	"strings"
)

func ValidateSearchTerm(term string) bool {
	if len(term) == 0 {
		return false
	}
	split := strings.Split(term, "/")
	if len(split) < 2 {
		return ValidateIdentifier(split[0])
	}
	organization := split[0]
	if !ValidateIdentifier(organization) {
		return false
	}
	repository := split[1]
	return ValidateIdentifier(repository)
}

func ValidateIdentifier(identifier string) bool {
	ok, _ := regexp.MatchString("^[A-Za-z0-9-]*$", identifier)
	return ok
}
