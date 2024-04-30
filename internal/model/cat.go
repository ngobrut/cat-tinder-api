package model

import (
	"time"

	"github.com/google/uuid"
)

type CatRace string

const (
	Persian          CatRace = "Persian"
	MaineCoon        CatRace = "Maine Coon"
	Siamese          CatRace = "Siamese"
	Ragdoll          CatRace = "Ragdoll"
	Bengal           CatRace = "Bengal"
	Sphynx           CatRace = "Sphynx"
	BritishShorthair CatRace = "BritishShorthair"
	Abyssinian       CatRace = "Abyssinian"
	ScottishFold     CatRace = "Scottish Fold"
	Birman           CatRace = "Birman"
)

var CatRaces = []string{
	string(Persian), string(MaineCoon), string(Siamese), string(Ragdoll), string(Bengal),
	string(Sphynx), string(BritishShorthair), string(Abyssinian), string(ScottishFold), string(Birman),
}

type CatSex string

const (
	Male   CatSex = "male"
	Female CatSex = "female"
)

var CatSexs = []string{
	string(Male), string(Female),
}

type Cat struct {
	CatID       uuid.UUID  `json:"cat_id" db:"cat_id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	Name        string     `json:"name" db:"name"`
	Race        CatRace    `json:"race" db:"race"`
	Sex         CatSex     `json:"sex" db:"sex"`
	AgeInMonth  int        `json:"age_in_month" db:"age_in_month"`
	Description string     `json:"description" db:"description"`
	HasMatched  *bool      `json:"has_matched" db:"has_matched"`
	ImageUrl    []string   `json:"image_url" db:"image_url"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
