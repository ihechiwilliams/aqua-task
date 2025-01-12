package constants

import "errors"

// ErrorCode ENUM(
//
// api_error
// validation_error,
//
// )
//
//go:generate go run github.com/abice/go-enum@v0.5.5
type ErrorCode string

var ErrRecordAlreadyExists = errors.New("record already exists")
