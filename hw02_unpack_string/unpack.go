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

	re_replace := regexp.MustCompile(`\D{1}0`)
	clearedStrToUnpack := re_replace.ReplaceAllString(strToUnpack, "")

	var builder strings.Builder
	var cached_rune rune
	for _, current_rune := range clearedStrToUnpack {
		decoded_rune, err := strconv.Atoi(string(current_rune))
		if err != nil {
			builder.WriteRune(current_rune)
		} else {
			builder.WriteString(strings.Repeat(string(cached_rune), decoded_rune-1))
		}
		cached_rune = current_rune
	}

	return builder.String(), nil
}
