package usecase

import (
	"database/sql"

	"github.com/ngobrut/cat-tinder-api/config"
	"github.com/ngobrut/cat-tinder-api/internal/repository"
)

type Usecase struct {
	cnf  *config.Config
	db   *sql.DB
	repo repository.IFaceRepository
}

func New(cnf *config.Config, db *sql.DB, repo repository.IFaceRepository) IFaceUsecase {
	return &Usecase{
		cnf:  cnf,
		db:   db,
		repo: repo,
	}
}
