package domain

import "errors"

var (
	ErrInvalidBalance  = errors.New("invalid balance")
	ErrBalanceNotFound = errors.New("balance not found")
	ErrInvalidData     = errors.New("invalid data")
	ErrInternal        = errors.New("internal server error")
	ErrEventNotFound   = errors.New("event not found")
)
