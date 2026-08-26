package domain

import "errors"

var (
	ErrOutOfStock      = errors.New("product out of stock")
	ErrProductNotFound = errors.New("product not found")
	ErrInternal        = errors.New("internal server error")
)
