package model

import (
	"time"

	"github.com/google/uuid"
)

type CatMatch struct {
	ID            uuid.UUID `json:"id" db:"id"`
	IssuerUserID  uuid.UUID `json:"issuer_user_id" db:"issuer_user_id"`
	IssuerCatID   uuid.UUID `json:"issuer_cat_id" db:"issuer_cat_id"`
	ReceiverCatID uuid.UUID `json:"receiver_cat_id" db:"receiver_cat_id"`
	Message       string    `json:"message" db:"message"`
	IsApproved    *bool     `json:"is_approved" db:"is_approved"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}