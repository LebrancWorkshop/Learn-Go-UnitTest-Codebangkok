package services

import "errors"

var (
	ErrNotValid = errors.New("amount is not valid")
	ErrRepository = errors.New("repository error")
)