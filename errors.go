package toolkit

import (
	"errors"
)

// Error represents error with http status code.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	// ErrUnauthorized ...
	ErrUnauthorized = errors.New("unauthorized")
)
