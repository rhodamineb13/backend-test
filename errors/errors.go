package customerrors

import (
	"fmt"
	"net/http"
)

type errors struct {
	statusCode uint
	message    string
}

func (e errors) Error() string {
	return fmt.Sprintf("Process returned with status code %d: %s", e.statusCode, e.message)
}

func newError(statusCode uint, message string) func(InnerError error) error {
	return func(InnerError error) error {
		return errors{
			statusCode: statusCode,
			message:    fmt.Sprintf("%s: %s", message, InnerError.Error()),
		}
	}
}

var (
	ErrUnexpected = newError(http.StatusInternalServerError, "unexpected error from database")
	ErrNotFound   = newError(http.StatusNotFound, "data not found")
)
