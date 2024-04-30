package custom_validator

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"

	"github.com/ngobrut/cat-tinder-api/internal/model"
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
			HTTPCode: http.StatusBadRequest,
			Message:  constant.ErrorMessageMap[http.StatusBadRequest],
		})

		return err
	}

	validate := validator.New()
	validate.RegisterValidation("catRace", validateCatRace)
	validate.RegisterValidation("catSex", validateCatSex)
	validate.RegisterValidation("imageUrls", validateImageUrls)
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
		case "catRace":
			message = fmt.Sprintf("%s must be one of %s", field.Field(), strings.Join(model.CatRaces, ", "))
		case "catSex":
			message = fmt.Sprintf("%s must be one of %s", field.Field(), strings.Join(model.CatSexs, ", "))
		case "imageUrls":
			message = fmt.Sprintf("%s must be greater than 1 and should be url", field.Field())
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

func validateCatRace(fl validator.FieldLevel) bool {
	race := model.CatRace(fl.Field().String())
	switch race {
	case model.Persian, model.MaineCoon, model.Siamese, model.Ragdoll, model.Bengal, model.Sphynx, model.BritishShorthair, model.Abyssinian, model.ScottishFold, model.Birman:
		return true
	default:
		return false
	}
}

func validateCatSex(fl validator.FieldLevel) bool {
	sex := model.CatSex(fl.Field().String())
	switch sex {
	case model.Male, model.Female:
		return true
	default:
		return false
	}
}

func validateImageUrls(fl validator.FieldLevel) bool {
	urls, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}
	if len(urls) < 1 {
		return false
	}

	for _, u := range urls {
		if u == "" || !isValidURL(u) {
			return false
		}
	}

	return true
}

func isValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}
