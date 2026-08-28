package domain

import "errors"

var (
	ErrInvalidBalance       = errors.New("invalid balance")
	ErrPaymentAlreadyExists = errors.New("payment already exists")
	ErrInvalidData          = errors.New("invalid data")
	ErrInternal             = errors.New("internal server error")
	ErrEventNotFound        = errors.New("event not found")
)
