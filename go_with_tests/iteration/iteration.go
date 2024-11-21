package iteration

import "strings"

// Constructs a string that repeats 'character' repeatCount times
// TODO: Use a string builder to reduce memory allocation
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
