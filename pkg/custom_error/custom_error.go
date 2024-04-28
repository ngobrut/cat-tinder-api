package custom_error

import (
	"net/http"

	"github.com/ngobrut/cat-tinder-api/pkg/constant"
)

type CustomError struct {
	ErrorContext *ErrorContext
}

type ErrorContext struct {
	HTTPCode int
	Message  string
}

func (c *CustomError) Error() string {
	if c.ErrorContext.HTTPCode == 0 {
		c.ErrorContext.HTTPCode = http.StatusInternalServerError
	}

	return constant.ErrorMessageMap[http.StatusInternalServerError]
}

func SetCustomError(errContext *ErrorContext) *CustomError {
	return &CustomError{
		ErrorContext: errContext,
	}
}
