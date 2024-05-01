package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/http/response"
)

type IFaceUsecase interface {
	// auth
	Register(req *request.Register) (*response.Register, error)
	Login(req *request.Login) (*response.Login, error)
	GetProfile(userID uuid.UUID) (*response.Profile, error)

	// manageCat
	CreateCat(req *request.CreateCat) (*response.CreateCat, error)
	GetListCat(params *request.ListCatQuery) ([]*response.CatResponse, error)
	UpdateCat(c *fiber.Ctx, req *request.UpdateCat) error
	DeleteCat(c *fiber.Ctx, catID uuid.UUID) error

	// cat match
	CreateCatMatch(c *fiber.Ctx, req *request.CreateCatMatch) error
	GetListCatMatch(params *request.ListCatMatchQuery) ([]*response.CatMatchResponse, error)
	ApproveCatMatch(c *fiber.Ctx, req *request.ApproveCatMatch) error
}
