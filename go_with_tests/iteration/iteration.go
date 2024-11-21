package iteration

import "strings"

// Constructs a string that repeats 'character' repeatCount times
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
