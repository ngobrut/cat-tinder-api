package constant

import "net/http"

// default error Code
const (
	DefaultUnhandledError = 1000 + iota
	DefaultNotFoundError
	DefaultBadRequestError
	DefaultUnauthorizedError
	DefaultDuplicateDataError
)

var ErrorMessageMap = map[int]string{
	http.StatusInternalServerError: "something went wrong with our side, please try again",
	http.StatusNotFound:            "data not found",
	http.StatusUnauthorized:        "you are not authorized to access this api",
	http.StatusConflict:            "duplicated data error",
	http.StatusUnprocessableEntity: "please check your body request",
	http.StatusBadRequest:          "request doesn't pass validation",
}
