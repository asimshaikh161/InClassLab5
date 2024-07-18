package main

import (
	"strings"
	"unicode"
)

// CamelToSnake converts a camelCase string into a snake_case string.
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			// Add an underscore before uppercase letters except for the first character
			if i > 0 {
				result.WriteRune('_')
			}
			// Convert uppercase letter to lowercase
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func main() {
	input := "thisIsCamelCase"
	output := CamelToSnake(input)
	println(output)
}
