package services

import "errors"

const (
	errDNAEmptyMessage = "DNA must be pass to function"
	errInvalidMatrix = "Matrix must be square. ex: 5x5 or 4x4"
	errInvalidCharDNA = "Invalid char in the matrix, only permitted: G, T, A, C."
)

var (
	ErrDNAEmpty = errors.New(errDNAEmptyMessage)
	ErrInvalidMatrix = errors.New(errInvalidMatrix)
	ErrInvalidCharMatrix = errors.New(errInvalidCharDNA)
)
