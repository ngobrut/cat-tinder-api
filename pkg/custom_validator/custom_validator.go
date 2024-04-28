package custom_validator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

type ValidatorError struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	Messages []string `json:"messages"`
}

func (o ValidatorError) Error() string {
	return "validate.request"
}

func ValidateStruct(c *fiber.Ctx, data interface{}) error {
	err := c.BodyParser(data)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusUnprocessableEntity,
			Message:  constant.ErrorMessageMap[http.StatusUnprocessableEntity],
		})

		return err
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err == nil {
		return nil
	}

	var message string
	var messages = make([]string, 0)

	for _, field := range err.(validator.ValidationErrors) {

		switch field.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", field.Field())
		case "email":
			message = fmt.Sprintf("%s is not in the correct email format", field.Field())
		case "min":
			message = fmt.Sprintf("%s must be %s character minimal", field.Field(), field.Param())
		case "max":
			message = fmt.Sprintf("%s cannot be more than %s character", field.Field(), field.Param())
		case "gt":
			message = fmt.Sprintf("%s must be greater than %s", field.Field(), field.Param())
		}

		messages = append(messages, message)
	}

	err = ValidatorError{
		Code:     http.StatusBadRequest,
		Message:  message,
		Messages: messages,
	}

	return err
}
