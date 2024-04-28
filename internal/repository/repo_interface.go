package repository

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

type IFaceRepository interface {
	// user
	CreateUser(data *model.User) error
	FindOneUserByID(userID uuid.UUID) (*model.User, error)
	FindOneUserByEmail(email string) (*model.User, error)
}
