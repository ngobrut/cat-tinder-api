package request

import "github.com/google/uuid"

type ListCatMatchQuery struct {
	UserID uuid.UUID
}

type CreateCatMatch struct {
	MatchCatID uuid.UUID `json:"matchCatId" validate:"required"`
	UserCatID  uuid.UUID `json:"userCatId" validate:"required"`
	Message    string    `json:"message" validate:"required,min=1,max=120"`
	UserID     uuid.UUID `json:"-"`
}

type RejectCatMatch struct {
	MatchID uuid.UUID `json:"matchId"`
	UserID  uuid.UUID `json:"-"`
}

type ApproveCatMatch struct {
	MatchID uuid.UUID `json:"matchId" validate:"required"`
	UserId  uuid.UUID `json:"-"`
}
