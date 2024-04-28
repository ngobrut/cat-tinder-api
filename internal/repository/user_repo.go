package repository

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/model"
	"github.com/ngobrut/cat-tinder-api/pkg/constant"
	"github.com/ngobrut/cat-tinder-api/pkg/custom_error"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(data *model.User) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	args := []interface{}{
		data.UserID,
		data.Name,
		data.Email,
		data.Password,
		timestamp,
		timestamp,
	}

	_, err := r.db.Exec("INSERT INTO users(user_id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", args...)
	if err != nil {
		if IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.ErrorMessageMap[http.StatusConflict],
			})
		}

		return err
	}

	return nil
}

// FindOneUserByID implements IFaceRepository.
func (r *Repository) FindOneUserByID(userID uuid.UUID) (*model.User, error) {
	res := &model.User{}

	if err := r.db.
		QueryRow("SELECT * FROM users WHERE user_id = $1 AND deleted_at IS NULL", userID).
		Scan(
			&res.UserID,
			&res.Name,
			&res.Email,
			&res.Password,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		); err != nil {
		return nil, err
	}

	return res, nil
}

// FindOneUser implements IFaceRepository.
func (r *Repository) FindOneUserByEmail(email string) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow("SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL", email).
		Scan(
			&res.UserID,
			&res.Name,
			&res.Email,
			&res.Password,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	if err != nil {
		return nil, err
	}

	return res, nil
}
