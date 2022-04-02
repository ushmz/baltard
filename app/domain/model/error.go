package model

import "github.com/pkg/errors"

var (
	ErrBadRequest          = errors.New("Invalid request")
	ErrNoSuchData          = errors.New("Requested data is not found")
	ErrInternalServerError = errors.New("Internal server error")
	ErrNilReciever         = errors.New("Reciever is nil")
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
