package helpers

import (
	"regexp"
)

var toEscape = regexp.MustCompile("([*_`[])")

func EscapeMarkdown(input string) string {
	var final []rune
	for _, c := range input {
		switch c {
		case '*', '_', '`', '[':
			final = append(final, '\\', c)
		default:
			final = append(final, c)
		}
	}
	return string(final)
}
