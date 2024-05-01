package response

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_validator"
)

type JsonResponse struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    *JsonMetadata  `json:"meta"`
	Error   *ErrorResponse `json:"error"`
}

type JsonMetadata struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func OK(c *fiber.Ctx, data interface{}, code int, message string) error {
	return c.Status(code).JSON(JsonResponse{
		Message: message,
		Data:    data,
	})
}

func Paging(c *fiber.Ctx, data interface{}, total int, limit int, offset int, message string) error {
	var metadata *JsonMetadata

	return c.Status(http.StatusOK).JSON(JsonResponse{
		Message: message,
		Data:    data,
		Meta:    metadata,
	})
}

func Error(c *fiber.Ctx, err error) error {
	v, isValidationErr := err.(custom_validator.ValidatorError)
	if isValidationErr {
		return c.Status(http.StatusBadRequest).JSON(JsonResponse{
			Message: "validation-error",
			Error: &ErrorResponse{
				Code:    v.Code,
				Message: v.Message,
			},
		})
	}

	e, isCustomErr := err.(*custom_error.CustomError)
	if !isCustomErr {
		fmt.Printf("[unhandled-error]\n%v\n", fmt.Sprint(err))

		return c.Status(http.StatusInternalServerError).JSON(JsonResponse{
			Message: constant.ErrorMessageMap[http.StatusInternalServerError],
			Error: &ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: constant.ErrorMessageMap[http.StatusInternalServerError],
			},
		})
	}

	code := http.StatusInternalServerError
	message := constant.ErrorMessageMap[constant.DefaultUnhandledError]

	if e.ErrorContext != nil && e.ErrorContext.HTTPCode > 0 {
		code = e.ErrorContext.HTTPCode
		message = constant.ErrorMessageMap[code]

		if e.ErrorContext.Message != "" {
			message = e.ErrorContext.Message
		}
	}

	return c.Status(code).JSON(JsonResponse{
		Message: "error:bad-request",
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
		},
	})
}
