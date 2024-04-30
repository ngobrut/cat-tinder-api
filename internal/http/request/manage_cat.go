package request

import (
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

type CreateCat struct {
	Name        string        `json:"name" validate:"required,min=1,max=30"`
	Race        model.CatRace `json:"race" validate:"required,catRace"`
	Sex         model.CatSex  `json:"sex" validate:"required,catSex"`
	AgeInMonth  int           `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string        `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string      `json:"imageUrls" validate:"required,imageUrls"`
}

type UpdateCat struct {
	Name        string        `json:"name" validate:"required,min=1,max=30"`
	Race        model.CatRace `json:"race" validate:"required,catRace"`
	Sex         model.CatSex  `json:"sex" validate:"required,catSex"`
	AgeInMonth  int           `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string        `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string      `json:"imageUrls" validate:"required,imageUrls"`
}
