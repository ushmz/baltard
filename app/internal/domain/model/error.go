package model

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
