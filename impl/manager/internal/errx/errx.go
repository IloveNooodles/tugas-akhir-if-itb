package errx

import (
	"errors"
	"strings"
)

var (
	ErrInvalidPassword  = errors.New("auth: invalid password")
	ErrInvalidToken     = errors.New("auth: invalid token")
	ErrDuplicateValuePq = errors.New("pq: duplicate key value")
	ErrClusterDown      = errors.New("connect: connection refused")
)

func IsDuplicateDatabase(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), ErrDuplicateValuePq.Error())
}

func IsClusterDown(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), ErrClusterDown.Error())
}
