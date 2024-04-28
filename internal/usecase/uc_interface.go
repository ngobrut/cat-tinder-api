package usecase

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

type IFaceUsecase interface {
	// auth
	Register(req *request.Register) (*response.Register, error)
	Login(req *request.Login) (*response.Login, error)
	GetProfile(userID uuid.UUID) (*model.User, error)
}
