package model

import (
	"golang.org/x/xerrors"
)

var (
	// ErrNilReceiver : The receiver of called function is nil
	ErrNilReceiver = xerrors.New("Reciever is nil")

	// ErrBadRequest : HTTP request body or argument is invalid
	ErrBadRequest = xerrors.New("Invalid request")

	// ErrNoSuchData : Requested data is not found
	ErrNoSuchData = xerrors.New("Requested data is not found")

	// ErrInternal : Internal errors that don't have to tell users in detail
	ErrInternal = xerrors.New("Internal server error")
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
