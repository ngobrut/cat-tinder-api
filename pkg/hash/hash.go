package hash

import (
	"net/http"

	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"golang.org/x/crypto/bcrypt"
)

func Compare(hashed []byte, plain []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hashed, plain)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "wrong password",
		})

		return
	}

	return
}

func HashAndSalt(cost int, pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
