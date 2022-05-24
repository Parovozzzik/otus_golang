package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var prev rune
	var result strings.Builder
	var withSlash int
	for _, q := range text {
		switch {
		case unicode.IsDigit(q):
			if (prev == 0 || unicode.IsDigit(prev)) && withSlash == 0 {
				return "", ErrInvalidString
			}
			processDigit(&result, prev, q, &withSlash)
		case unicode.IsLetter(q):
			processLetter(&result, prev, q, &withSlash)
		case unicode.IsPunct(q) && string(q) == `\`:
			if string(prev) == `\` && withSlash == 0 {
				result.WriteRune(q)
				withSlash = 1
				continue
			}
			if withSlash == 1 {
				withSlash = 0
			}
		}
		prev = q
	}

	if result.String() == "" && text != "" {
		return "", ErrInvalidString
	}
	return result.String(), nil
}

func processDigit(result *strings.Builder, prev rune, q rune, withSlash *int) {
	if int(q-'0') > 0 {
		if string(prev) != `\` || *withSlash == 1 {
			repeat(result, prev, int(q-'0')-1, withSlash)
		} else {
			result.WriteRune(q)
			*withSlash = 1
		}
	} else {
		cut(result)
	}
}

func processLetter(result *strings.Builder, prev rune, q rune, withSlash *int) {
	if string(prev) == `\` && string(q) == `n` {
		result.WriteRune(prev)
		*withSlash = 1
	}
	result.WriteRune(q)
}

func repeat(result *strings.Builder, prev rune, count int, withSlash *int) {
	for i := 0; i < count; i++ {
		if *withSlash == 1 && string(prev) == `n` {
			result.WriteRune('\\')
		}
		result.WriteRune(prev)
	}
}

func cut(result *strings.Builder) {
	currentText := result.String()
	currentText = currentText[:len(currentText)-1]
	result.Reset()
	for _, q := range currentText {
		result.WriteRune(q)
	}
}
