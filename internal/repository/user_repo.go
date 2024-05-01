package repository

import (
	"github.com/google/uuid"
	"github.com/ngobrut/cat-tinder-api/internal/model"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(data *model.User) error {
	args := []interface{}{
		data.UserID,
		data.Name,
		data.Email,
		data.Password,
		data.CreatedAt.Format("2006-01-02 15:04:05"),
		data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	_, err := r.db.Exec("INSERT INTO users(user_id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", args...)
	if err != nil {
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

	if err := r.db.
		QueryRow("SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL", email).
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
