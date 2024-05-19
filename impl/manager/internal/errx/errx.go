package errx

import "errors"

var (
	ErrInvalidPassword = errors.New("auth: invalid password")
	ErrInvalidToken    = errors.New("auth: invalid token")
)
