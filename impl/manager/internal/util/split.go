package util

import (
	"errors"
	"strings"
)

var (
	ErrInvalidKeyVal = errors.New("validation: invalid key val")
)

func SplitByComma(input string) []string {
	return strings.Split(input, ",")
}

func SplitByEqual(input string) (key string, val string, err error) {
	res := strings.Split(input, "=")
	if len(res) != 2 {
		err = ErrInvalidKeyVal
		return
	}

	key = res[0]
	val = res[1]
	return
}
