package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
)

type Cat struct {
	CatID       uuid.UUID        `json:"cat_id" db:"cat_id"`
	UserID      uuid.UUID        `json:"user_id" db:"user_id"`
	Name        string           `json:"name" db:"name"`
	Race        constant.CatRace `json:"race" db:"race"`
	Sex         constant.CatSex  `json:"sex" db:"sex"`
	AgeInMonth  int              `json:"age_in_month" db:"age_in_month"`
	Description string           `json:"description" db:"description"`
	HasMatched  *bool            `json:"has_matched" db:"has_matched"`
	ImageUrl    []string         `json:"image_url" db:"image_url"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time       `json:"deleted_at" db:"deleted_at"`
}
