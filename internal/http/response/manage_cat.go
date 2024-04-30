package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

type CreateCat struct {
	CatID     uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
}

type GetCats struct {
	CatID       uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Race        model.CatRace `json:"race"`
	Sex         model.CatSex  `json:"sex"`
	AgeInMonth  int           `json:"ageInMonth"`
	ImageUrls   []string      `json:"imageUrls"`
	Description string        `json:"description"`
	HasMatched  bool          `json:"hasMatched"`
	CreatedAt   time.Time     `json:"createdAt"`
}
