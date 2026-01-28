package helpers

import (
	"regexp"
	"strings"
)

func ToSnakeCase(s string) string {
	// 1. Handle transitions from acronyms to words (e.g., "HTTP" -> "HTTP_")
	// Matches uppercase letters followed by an uppercase and then a lowercase
	var matchAcronym = regexp.MustCompile("([A-Z]+)([A-Z][a-z])")

	// 2. Handle transitions from lowercase/numbers to uppercase (e.g., "newCar" -> "new_Car")
	var matchFirstCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	res := matchAcronym.ReplaceAllString(s, "${1}_${2}")
	res = matchFirstCap.ReplaceAllString(res, "${1}_${2}")

	return strings.ToLower(res)
}
