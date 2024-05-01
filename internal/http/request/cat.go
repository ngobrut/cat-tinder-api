package request

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
)

type ListCatQuery struct {
	Limit      int              `query:"limit"`
	Offset     int              `query:"offset"`
	ID         string           `query:"id"`
	Race       constant.CatRace `query:"race"`
	Sex        constant.CatSex  `query:"sex"`
	HasMatched string           `query:"hasMatched"`
	AgeInMonth string           `query:"ageInMonth"`
	Owned      string           `query:"owned"`
	Search     string           `query:"search"`
	UserID     uuid.UUID        `query:"-"`
}

type CreateCat struct {
	Name        string           `json:"name" validate:"required,min=1,max=30"`
	Race        constant.CatRace `json:"race" validate:"required,catRace"`
	Sex         constant.CatSex  `json:"sex" validate:"required,catSex"`
	AgeInMonth  int              `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string           `json:"description" validate:"required,min=1,max=200"`
	ImageURLs   []string         `json:"imageUrls" validate:"required,gt=0,dive,url"`
	UserID      uuid.UUID        `json:"-"`
}

type UpdateCat struct {
	Name        string           `json:"name" validate:"required,min=1,max=30"`
	Race        constant.CatRace `json:"race" validate:"required,catRace"`
	Sex         constant.CatSex  `json:"sex" validate:"required,catSex"`
	AgeInMonth  int              `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string           `json:"description" validate:"required,min=1,max=200"`
	ImageURLs   []string         `json:"imageUrls" validate:"required,required,gt=0,dive,url"`
	CatID       uuid.UUID        `json:"-"`
	UserID      uuid.UUID        `json:"-"`
}
