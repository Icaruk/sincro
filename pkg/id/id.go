package id

import (
	"regexp"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Generate(name string) string {

	if len(name) > 8 {
		name = name[:8]
	}

	id, err := gonanoid.Generate(CHARACTERS, 16)
	if err != nil {
		return ""
	}

	return name + "_" + id

}

func Validate(id string) bool {

	// Split with _
	split := strings.Split(id, "_")

	if len(split) != 2 {
		return false
	}

	name := split[0]
	subid := split[1]

	// Validate name fragment
	if len(name) > 8 || len(name) == 0 {
		return false
	}
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	isNameValid := r.MatchString(name)

	if !isNameValid {
		return false
	}

	// Validate subid fragment
	if len(subid) != 16 || len(subid) == 0 {
		return false
	}

	isSubidValid := r.MatchString(subid)
	if !isSubidValid {
		return false
	}

	return true

}
