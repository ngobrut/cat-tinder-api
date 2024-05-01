package custom_validator

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

type ValidatorError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterValidation("catRace", validateCatRace)
	validate.RegisterValidation("catSex", validateCatSex)
	validate.RegisterValidation("imageUrls", validateImageUrls)

	err = validate.Struct(data)
	if err == nil {
		return nil
	}

	var message string
	for _, field := range err.(validator.ValidationErrors) {
		message = field.Translate(trans)

		switch field.Tag() {
		case "catRace":
			message = fmt.Sprintf("%s must be one of [%s]", field.Field(), strings.Join(constant.CatRaces, ", "))
		case "catSex":
			message = fmt.Sprintf("%s must be one of [%s]", field.Field(), strings.Join(constant.CatSexes, ", "))
		case "imageUrls":
			message = fmt.Sprintf("%s must be greater than 1 and should be url", field.Field())
		}
	}

	err = ValidatorError{
		Code:    http.StatusBadRequest,
		Message: message,
	}

	return err
}

func validateCatRace(fl validator.FieldLevel) bool {
	return constant.ValidCatRace[fl.Field().String()]
}

func validateCatSex(fl validator.FieldLevel) bool {
	return constant.ValidCatSex[fl.Field().String()]
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
	if err != nil {
		return false
	}

	return u.Scheme != "" && u.Host != ""
}
