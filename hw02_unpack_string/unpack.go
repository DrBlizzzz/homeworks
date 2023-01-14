package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(strToUnpack string) (string, error) {
	rules := []*regexp.Regexp{
		regexp.MustCompile(`\d{2}`),
		regexp.MustCompile(`^\d{1}`),
	}

	for _, rule := range rules {
		if rule.FindString(strToUnpack) != "" {
			return "", ErrInvalidString
		}
	}

	reReplace := regexp.MustCompile(`\D{1}0`)
	clearedStrToUnpack := reReplace.ReplaceAllString(strToUnpack, "")

	var builder strings.Builder
	var cachedRune rune
	for _, currentRune := range clearedStrToUnpack {
		decodedRune, err := strconv.Atoi(string(currentRune))
		if err != nil {
			builder.WriteRune(currentRune)
		} else {
			builder.WriteString(strings.Repeat(string(cachedRune), decodedRune-1))
		}
		cachedRune = currentRune
	}

	return builder.String(), nil
}
