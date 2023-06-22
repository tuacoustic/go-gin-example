package common

import (
	"strings"
	"unicode"
)

func CamelToSnake(camelCase string) string {
	var sb strings.Builder
	sb.Grow(len(camelCase) + 5) // Preallocate enough space to avoid reallocation

	for i, r := range camelCase {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
