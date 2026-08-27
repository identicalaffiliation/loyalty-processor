package domain

import "errors"

var (
	ErrInvalidData     = errors.New("invalid data")
	ErrInternal        = errors.New("internal server error")
	ErrOutOfStock      = errors.New("product out of stock")
	ErrProductNotFound = errors.New("product not found")
	ErrEventNotFound   = errors.New("event not found")
)
