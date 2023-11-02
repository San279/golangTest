package services

import "errors"

var (
	ErrZeroAmount = errors.New("amount could not be zero")
	ErrRepository = errors.New("repository error")
)
