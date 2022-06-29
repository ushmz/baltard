package model

import (
	"errors"
)

var (
	// ErrDBOperationFailed : DB operation failed
	ErrDBOperationFailed = errors.New("DB operation failed")

	// ErrDBConnectionFailed : DB connection fauled
	ErrDBConnectionFailed = errors.New("DB connection fauled")

	// ErrNilReceiver : The receiver of called function is nil
	ErrNilReceiver = errors.New("Called with nil reciever")

	// ErrBadRequest : HTTP request body or argument is invalid
	ErrBadRequest = errors.New("Invalid request")

	// ErrNoSuchData : Requested data is not found
	ErrNoSuchData = errors.New("Requested data is not found")

	// ErrInternal : Internal errors that don't have to tell users in detail
	ErrInternal = errors.New("Internal server error")
)

// ErrorMessage : Struct for error message response.
type ErrorMessage struct {
	// Message : Error message string.
	Message string `json:"message"`
}

// ErrWithMessage : Wrap errors with response messsage
type ErrWithMessage struct {
	error
	Why string
}

// NewErrWithMessage : Return new ErrWithMessage
func NewErrWithMessage(err error, why string) error {
	return ErrWithMessage{err, why}
}

func (e *ErrWithMessage) Unwrap() error {
	return e.error
}
