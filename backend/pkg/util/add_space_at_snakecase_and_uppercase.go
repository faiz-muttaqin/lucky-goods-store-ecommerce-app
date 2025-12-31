package util

import (
	"strings"
	"unicode"
)

func AddSpaceAtSnakeCaseAndUppercase(s string) string {
	if s == "" {
		return s
	}

	var result []rune
	capitalizeNext := true // First character should be uppercase

	for i, r := range s {
		if r == '_' {
			// Replace underscore with space
			result = append(result, ' ')
			capitalizeNext = true // Next character should be uppercase
		} else if unicode.IsUpper(r) && i != 0 && !capitalizeNext {
			// Add space before uppercase letter (but not at start or after underscore)
			result = append(result, ' ')
			result = append(result, r)
			capitalizeNext = false
		} else if capitalizeNext {
			// Capitalize this character
			result = append(result, unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			// Keep character as-is
			result = append(result, r)
		}
	}

	return strings.TrimSpace(string(result))
}
