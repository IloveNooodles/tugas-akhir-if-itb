package errx

import (
	"errors"
	"strings"
)

var (
	ErrInvalidPassword  = errors.New("auth: invalid password")
	ErrInvalidToken     = errors.New("auth: invalid token")
	ErrDuplicateValuePq = errors.New("pq: duplicate key value")
	ErrClusterDown      = errors.New("server is down, please try again later")
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

	return strings.Contains(err.Error(), "connect: connection refused")
}

func IsNodeNotFound(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "not found")
}
