package errors

import "errors"

var (
	ErrNilLogger  = errors.New("logger not initialized")
	ErrNilStorage = errors.New("storage not initialized")
)
