package response

import (
	"time"

	"github.com/google/uuid"
)

type IssuedBy struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CatMatchResponse struct {
	Id             uuid.UUID    `json:"id"`
	IssuedBy       *IssuedBy    `json:"issuedBy"`
	MatchCatDetail *CatResponse `json:"matchCatDetail"`
	UserCatDetail  *CatResponse `json:"userCatDetail"`
	Message        string       `json:"message"`
	CreatedAt      time.Time    `json:"createdAt"`
}
