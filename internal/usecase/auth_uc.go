package usecase

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
	"github.com/ngobrut/cat-tinder-api/pkg/hash"
	"github.com/ngobrut/cat-tinder-api/pkg/jwt"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(req *request.Register) (*response.Register, error) {
	existing, err := u.repo.FindOneUserByEmail(req.Email)
	if err != nil && !repository.IsRecordNotFound(err) {
		return nil, err
	}

	if existing != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusConflict,
			Message:  "email has been used",
		})

		return nil, err
	}

	pwd, err := hash.HashAndSalt(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UserID:    uuid.New(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  pwd,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.repo.CreateUser(user)
	if err != nil {
		if repository.IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  "email has been used",
			})
		}

		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: user.UserID.String(),
	}

	token, err := jwt.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.Register{
		Email:       req.Email,
		Name:        req.Name,
		AccessToken: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(req *request.Login) (*response.Login, error) {
	user, err := u.repo.FindOneUserByEmail(req.Email)
	if err != nil && !repository.IsRecordNotFound(err) {
		return nil, err
	}

	if user == nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "no user with requested email found",
		})

		return nil, err
	}

	err = hash.Compare([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomClaims{
		UserID: user.UserID.String(),
	}

	token, err := jwt.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.Login{
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: token,
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(userID uuid.UUID) (*response.Profile, error) {
	user, err := u.repo.FindOneUserByID(userID)
	if err != nil {
		return nil, err
	}

	res := &response.Profile{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}
