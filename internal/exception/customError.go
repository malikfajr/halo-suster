package exception

import "net/http"

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Error implements error.
func (*CustomError) Error() string {
	panic("unimplemented")
}

func NewBadRequest(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewConflict(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

func NewInternalError(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewNotFound(message string) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}
