package model

import "github.com/pkg/errors"

var (
	// ErrBadRequest : HTTP request body or argument is invalid
	ErrBadRequest = errors.New("Invalid request")

	// ErrNoSuchData : Requested data is not found
	ErrNoSuchData = errors.New("Requested data is not found")

	// ErrInternalServerError : Internal errors that don't have to tell users in detail
	ErrInternalServerError = errors.New("Internal server error")

	// ErrNilReceiver : The receiver of called function is nil
	ErrNilReceiver = errors.New("Reciever is nil")
)

// ErrorMessage : Struct for error message response.
type ErrorMessage struct {
	// Message : Error message string.
	Message string `json:"message"`
}

// NoSuchDataError : Wrap no such data error like `sql.ErrNoRows`
type NoSuchDataError struct{}

func (e NoSuchDataError) Error() string {
	return "No such data"
}
