package validation

import (
	"errors"
	"regexp"
)

func ValidateProjectName(name string) (isValid bool, error error) {
	if len(name) > 8 || len(name) == 0 {
		return false, errors.New("name must have 1-8 characters")
	}
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	isNameValid := r.MatchString(name)
	if !isNameValid {
		return false, errors.New("name must only have letters and numbers")
	}
	return true, nil
}
