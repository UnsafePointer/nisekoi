package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateSearchTerm(term string) (string, string, error) {
	if len(term) == 0 {
		return "", "", errors.New("Search term empty")
	}
	split := strings.Split(term, "/")
	if len(split) < 2 {
		if err := ValidateIdentifier(split[0]); err != nil {
			return "", "", err
		}
		return split[0], "", nil
	}
	owner := split[0]
	if err := ValidateIdentifier(owner); err != nil {
		return "", "", err
	}
	repository := split[1]
	if err := ValidateIdentifier(repository); err != nil {
		return "", "", err
	}
	return owner, repository, ValidateIdentifier(repository)
}

func ValidateIdentifier(identifier string) error {
	ok, error := regexp.MatchString("^[A-Za-z0-9-]*$", identifier)
	if !ok || error != nil {
		return errors.New("The provided search term contains invalid characters")
	}
	return nil
}
