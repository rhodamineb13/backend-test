package customerrors

import (
	"fmt"
	"net/http"
)

type Errors struct {
	StatusCode uint
	Message    string
}

func (e Errors) Error() string {
	return fmt.Sprintf("Process returned with status code %d: %s", e.StatusCode, e.Message)
}

func newError(statusCode uint, message string) func(InnerError error) Errors {
	return func(InnerError error) Errors {
		return Errors{
			StatusCode: statusCode,
			Message:    fmt.Sprintf("%s: %s", message, InnerError.Error()),
		}
	}
}

var (
	ErrUnexpected = newError(http.StatusInternalServerError, "unexpected error from database")
	ErrNotFound   = newError(http.StatusNotFound, "data not found")
	ErrBadRequest = newError(http.StatusBadRequest, "bad request from client")
)
