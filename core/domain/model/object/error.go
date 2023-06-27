package domain

import (
	"errors"
)

var (
	ErrNil                      error = nil
	ErrInvalidInput             error = errors.New("invalid input")
	ErrCommandIsNotShellCommand error = errors.New("command is not shell command")
	ErrPipeIsNotSupported       error = errors.New("pipe is not supported")
)
