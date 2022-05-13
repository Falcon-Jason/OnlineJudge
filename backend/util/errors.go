package util

import "errors"

var (
	ErrInvalidRequest  = errors.New("invalid request")
	ErrAuthorization   = errors.New("authorization failed")
	ErrNothingModified = errors.New("nothing modified")
)
