package http_errors

import "errors"

var (
	CodeExpired = errors.New("code is expired")
)
