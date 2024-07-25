package utils

import (
	"strings"
)

// Checks if the first string (needle) is in the second one (haystack)
// starting from the beginning
func Match(needle string, haystack string) bool {
	needle = strings.ToLower(needle)
	haystack = strings.ToLower(haystack)

	if len(needle) > len(haystack) {
		return needle[:len(haystack)] == haystack
	}

	return haystack[:len(needle)] == needle
}
