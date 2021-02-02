package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder
	var prev rune
	var prevS string

	for _, cur := range str {
		prevS = string(prev)

		if unicode.IsDigit(cur) &&
			(unicode.IsDigit(prev) || prev == 0 || prevS == " ") {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(cur) {
			num, _ := strconv.Atoi(string(cur))

			prevS = strings.Repeat(prevS, num)
		}

		if isValid(prev) {
			res.WriteString(prevS)
		}

		prev = cur
	}

	if isValid(prev) {
		res.WriteString(string(prev))
	}

	return res.String(), nil
}

func isValid(prev rune) bool {
	if unicode.IsDigit(prev) || prev == 0 {
		return false
	}

	return true
}
