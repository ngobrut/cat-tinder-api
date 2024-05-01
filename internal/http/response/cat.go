package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
)

type CreateCat struct {
	CatID     uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
}

type CatResponse struct {
	CatID       uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Race        constant.CatRace `json:"race"`
	Sex         constant.CatSex  `json:"sex"`
	AgeInMonth  int              `json:"ageInMonth"`
	ImageUrls   []string         `json:"imageUrls"`
	Description string           `json:"description"`
	HasMatched  bool             `json:"hasMatched"`
	CreatedAt   time.Time        `json:"createdAt"`
}
