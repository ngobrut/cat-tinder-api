package repository

import (
	"database/sql"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) IFaceRepository {
	return &Repository{
		db: db,
	}
}

func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}

func IsRecordNotFound(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}
