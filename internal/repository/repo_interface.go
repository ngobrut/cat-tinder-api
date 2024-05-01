package repository

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/http/request"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

type IFaceRepository interface {
	// user
	CreateUser(data *model.User) error
	FindOneUserByID(userID uuid.UUID) (*model.User, error)
	FindOneUserByEmail(email string) (*model.User, error)

	// cat
	CreateCat(data *model.Cat) error
	FindCat(params *request.ListCatQuery) ([]*model.Cat, error)
	FindOneCatByID(catID uuid.UUID) (*model.Cat, error)
	CheckOwnCat(userID uuid.UUID, catID uuid.UUID) error
	UpdateCatByID(data map[string]interface{}, catID uuid.UUID) error
	DeleteCatByID(catID uuid.UUID) error

	// cat match
	CreateCatMatch(data *model.CatMatch) error
	CheckDuplicateMatchRequest(issuerCatID uuid.UUID, receiverCatID uuid.UUID) error
	FindOneCatMatchByCatID(catID uuid.UUID) (*model.CatMatch, error)
	FindCatMatch(params *request.ListCatMatchQuery) ([]*model.CatMatch, error)
	FindOneCatMatchByID(ID uuid.UUID) (*model.CatMatch, error)
	UpdateCatMatchByID(data map[string]interface{}, ID uuid.UUID) error
	DeleteCatMatchByID(ID uuid.UUID) error
}
