package logger

import "errors"

var (
	ErrInvalidFormat = errors.New("invalid format")
	ErrInvalidLevel  = errors.New("invalid logging level")
)
